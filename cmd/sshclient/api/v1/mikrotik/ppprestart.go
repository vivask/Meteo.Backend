package mikrotik

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/log"
	"net/http"
	"regexp"
	"strconv"
	s "strings"
	"time"

	SSH "meteo/cmd/sshclient/api/v1/internal"

	"github.com/gin-gonic/gin"
)

const (
	MAX_TIME = 6
	MIN_TIME = 1
)

func (p mikrotikAPI) CheckBYFLY(c *gin.Context) {

	err := p.checkPPPState()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SSHERR-1",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p mikrotikAPI) checkPPPState() error {

	bind := fmt.Sprintf("%s:%d", config.Default.SshClient.PPP.Host, config.Default.SshClient.PPP.Port)

	link, err := SSH.NewSSHLink(bind, config.Default.SshClient.PPP.User)
	if err != nil {
		return fmt.Errorf("can't create ssh link: %w", err)
	}
	defer link.Close()

	cmd := fmt.Sprintf("/interface pppoe-client monitor %s once", config.Default.SshClient.PPP.Interface)
	req, err := link.Exec(cmd, 1)
	if err != nil {
		return fmt.Errorf("ssh exec error: %w", err)
	}
	matched, _ := regexp.MatchString("status: connected", req)
	if !matched {
		return fmt.Errorf("ssh response error: %s", req)
	}
	log.Debugf("PPPOE status: %s", req)

	expr, err := extractExpr(req, "uptime: (.*)")
	if err != nil {
		return fmt.Errorf("extract expression error: %w", err)
	}
	ticks, err := getTicks(expr)
	if err != nil {
		return fmt.Errorf("get ticks error: %w", err)
	}
	hr := time.Now().Local().Add(ticks).Hour()

	if hr > MAX_TIME || hr < MIN_TIME {
		cmd = fmt.Sprintf("/system scheduler print from %s value-list", config.Default.SshClient.PPP.Script)
		answer := fmt.Sprintf("name: %s", config.Default.SshClient.PPP.Script)
		req, err := link.Exec(cmd, 1)
		if err != nil {
			return fmt.Errorf("ssh exec error: %w", err)
		}
		matched, _ := regexp.MatchString(answer, req)
		if !matched {
			return fmt.Errorf("ssh response error: %s", req)
		}
		log.Debugf("PPPScript: %s", req)
		expr, err = extractExpr(req, "start-date: (.*)")
		if err != nil {
			return fmt.Errorf("extract expression error: %w", err)
		}
		start_date, err := mikrotikDateToDate(expr)
		if err != nil {
			return fmt.Errorf("convert date error: %w", err)
		}
		date_restart := dateRestart(hr, ticks)
		days := diffDate(date_restart, start_date)
		if days > 1 {
			_, err = link.Exec("/log warning \"PPP restart scrip activated!\"", 1)
			if err != nil {
				return fmt.Errorf("ssh exec error: %w", err)
			}
			cmd = fmt.Sprintf("/system scheduler set %s start-date=%s", config.Default.SshClient.PPP.Script, dateToMikrotikDate(date_restart))
			_, err = link.Exec(cmd, 1)
			if err != nil {
				return fmt.Errorf("ssh exec error: %w", err)
			}
			log.Infof("Date ppp restart: %v", dateToMikrotikDate(date_restart))
		}
	}
	return nil
}

func getTicks(str string) (ticks time.Duration, err error) {

	var duration time.Duration
	var days int
	s := s.Split(str, "d")

	if len(s) > 1 {
		duration, err = time.ParseDuration(s[1])
		if err != nil {
			return 0, err
		}
		days, err = strconv.Atoi(s[0])
		if err != nil {
			return 0, err
		}
	} else {
		duration, err = time.ParseDuration(s[0])
		if err != nil {
			return 0, err
		}
		days = 0
	}

	return time.Duration(432000-duration.Seconds()-float64(days*24*60*60)) * time.Second, nil
}

func extractExpr(str, expr string) (res string, err error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return
	}
	updatestr := r.FindString(str)
	split := s.Split(updatestr, " ")[1]
	res = split[:len(split)-1]
	return
}

func mikrotikDateToDate(d string) (date time.Time, err error) {

	split := s.Split(d, "/")
	m := split[0]
	day := split[1]
	year := split[2]
	months := map[string]string{"jan": "01", "feb": "02", "mar": "03", "apr": "04", "may": "05", "jun": "06", "jul": "07", "aug": "08", "sep": "09", "oct": "10", "now": "11", "dec": "12"}
	month := months[m]
	date, err = time.Parse("2006-01-02", year+"-"+month+"-"+day)
	return
}

func dateToMikrotikDate(d time.Time) (date string) {

	months := map[int]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Now", 12: "Dec"}
	return months[int(d.Month())] + "/" + strconv.Itoa(d.Day()) + "/" + strconv.Itoa(d.Year())
}

func diffDate(d1 time.Time, d2 time.Time) (diff int64) {
	u1 := d1.Unix()
	u2 := d2.Unix()
	if u1 > u2 {
		diff = (u1 - u2) / 86400
	} else {
		diff = (u2 - u1) / 86400
	}
	return
}

func dateRestart(hour int, ticks time.Duration) (dr time.Time) {
	if hour >= MIN_TIME && hour <= MAX_TIME {
		dr = time.Now().Local().Add(ticks)
	} else {
		dr = time.Now().Local().AddDate(0, 0, 1)
	}
	return
}

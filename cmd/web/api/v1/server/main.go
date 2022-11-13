package server

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/kit"
	"meteo/internal/log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Services struct {
	ClusterService   bool
	MessangerService bool
	ServerService    bool
	Esp32Service     bool
	ProxyService     bool
	SshclientService bool
	ScheduleService  bool
	WebService       bool
	GogsService      bool
	PostgresService  bool
}

func (p serverAPI) GetMainServicesState(c *gin.Context) {

	var state Services

	req, err := kit.GetMain("/cluster/health")
	state.ClusterService = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/messanger/health")
	state.MessangerService = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/server/health")
	state.ServerService = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/esp32/health")
	state.Esp32Service = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/proxy/health")
	state.ProxyService = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/sshclient/health")
	state.SshclientService = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/schedule/health")
	state.ScheduleService = (err == nil) && (len(req) == 0)

	req, err = kit.GetMain("/web/health")
	state.WebService = (err == nil) && (len(req) == 0)

	state.GogsService = raw_connect(config.Default.Client.Local, []string{"2222", "3000"})

	state.PostgresService = raw_connect(config.Default.Client.Local, []string{"5432"})

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": state})
}

func raw_connect(host string, ports []string) bool {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			log.Errorf("Connect to %s:%s fail: %v", host, port, err)
			return false
		}
		if conn != nil {
			defer conn.Close()
		}
	}
	return true
}

func (p serverAPI) RestartMainServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/main/restart/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StopMainServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/main/stop/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) StartMainServerCont(c *gin.Context) {

	url := fmt.Sprintf("/sshclient/server/main/start/%s", c.Param("id"))
	_, err := kit.PutMain(url, nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) MainReboot(c *gin.Context) {

	_, err := kit.PutMain("/sshclient/server/main/reboot", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (p serverAPI) MainShutdown(c *gin.Context) {

	_, err := kit.PutMain("/sshclient/server/main/shutdown", nil)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code":    http.StatusInternalServerError,
				"error":   "SERVER",
				"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

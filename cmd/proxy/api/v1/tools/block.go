package tools

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	s "strings"

	"meteo/internal/dto"
	"meteo/internal/entities"
	"meteo/internal/log"
	repo "meteo/internal/repo/proxy"
)

const blockListDir = "/var/lib/proxy/data"
const blockListFileName = "blocklist.txt"

type BlackList struct {
	data map[string]struct{}
}

func NewBlackList() (*BlackList, error) {
	if _, err := os.Stat(blockListDir); os.IsNotExist(err) {
		err := os.Mkdir(blockListDir, 0755)
		if err != nil {
			return nil, fmt.Errorf("can't create dir: %w", err)
		}
	}

	return &BlackList{
		data: make(map[string]struct{}),
	}, nil
}

func (b *BlackList) GetData() map[string]struct{} {
	return b.data
}

func (b *BlackList) SetData(m map[string]struct{}) {
	b.data = m
	log.Infof("Receiving and set %d block hosts", len(m))
}

func Split(r rune) bool {
	return r == ' ' || r == '\t'
}

func (b *BlackList) Count() int {
	return len(b.data)
}

func (b *BlackList) Add(host string) bool {

	if _, exist := b.data[host]; !exist {
		b.data[host] = struct{}{}
		return true
	}
	return false
}

func extractName(line string) (name string, good bool) {

	if len(line) == 0 || line[0:1] == "#" || line[0:1] == "!" || badHost(line) {
		return "", false
	}

	if len(line) > 2 && line[0:2] == "||" {
		name = line[2 : len(line)-1]
		if !s.HasSuffix(name, ".") {
			name += "."
		}

		return name, true
	}

	if len(line) > 10 && (line[0:9] == "127.0.0.1" || line[0:7] == "0.0.0.0") {
		split := s.FieldsFunc(line, Split)
		name = s.Trim(split[1], " ")
		if len(name) == 0 {
			return "", false
		}

		if !s.HasSuffix(name, ".") {
			name += "."
		}

		return name, true
	} else {
		return "", false
	}
}

func badHost(host string) bool {

	badhosts := []string{
		"localhost",
		"localhost.localdomain",
		"broadcasthost",
		"local",
		"ip6-localhost",
		"ip6-loopback",
		"ip6-localnet",
		"ip6-mcastprefix",
		"ip6-allnodes",
		"ip6-allrouters",
		"ip6-allhosts",
	}

	for _, bad := range badhosts {
		matched, _ := regexp.MatchString(bad, host)
		if matched {
			return true
		}
	}
	return false
}

func (b *BlackList) AddList(lines []string, idx int) (count int) {

	for _, line := range lines {
		name, ok := extractName(line)
		if ok && b.Add(name) {
			count++
		}
	}

	return
}

func (b *BlackList) loadFromDb(repo repo.ProxyService) (count int) {

	hosts, err := repo.GetAllBlockHosts(dto.Pageable{})
	if err != nil {
		log.Error(err)
	}
	for _, host := range *hosts {
		b.data[host.ID] = struct{}{}
		count++
	}
	return
}

func (b *BlackList) saveToDb(repo repo.ProxyService) error {
	err := repo.ClearBlocklist()
	if err != nil {
		return fmt.Errorf("clear blocklist error: %w", err)
	}

	for host := range b.data {
		err := repo.AddBlockHost(entities.Blocklist{ID: host})
		if err != nil {
			return fmt.Errorf("add blocklist error: %w", err)
		}
	}
	return nil
}

func (b *BlackList) loadFromFile() (count int, err error) {
	fName := fmt.Sprintf("%s/%s", blockListDir, blockListFileName)
	f, err := os.OpenFile(fName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, fmt.Errorf("open file error: %w", err)
	}
	defer f.Close()
	buf := new(s.Builder)
	io.Copy(buf, f)
	hosts := s.Split(buf.String(), "\n")
	for _, host := range hosts {
		if len(s.Trim(host, " ")) != 0 {
			b.data[host] = struct{}{}
			count++
		}
	}
	return
}

func (b *BlackList) SaveToFile() error {
	if len(b.data) == 0 {
		return errors.New("ad block list is empty")
	}

	fName := fmt.Sprintf("%s/%s", blockListDir, blockListFileName)
	f, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}
	defer f.Close()

	for k := range b.data {
		_, err = io.WriteString(f, fmt.Sprintf("%s\n", k))
		if err != nil {
			return fmt.Errorf("file write error: %w", err)
		}
	}
	return nil
}

func (b *BlackList) Contains(server string) bool {
	_, ok := b.data[server]
	return ok
}

func updateList(lists []string) (*BlackList, []bool) {

	loaded := make([]bool, len(lists))
	list, err := NewBlackList()
	if err != nil {
		log.Errorf("Internal error: %v", err)
		return list, loaded
	}

	for idx, v := range lists {
		resp, err := http.Get(v)
		if err != nil {
			log.Errorf("[black] Can't load: %s", v)
			log.Error(err)
			continue
		}

		if resp.StatusCode != 200 {
			log.Errorf("[black] Status code of %s!= 200", v)
			continue
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("[black] Can't read body of %s", v)
			continue
		}

		data2 := s.Split(string(data), "\n")
		cnt := list.AddList(data2, idx)

		if cnt == 0 {
			loaded[idx] = false
			log.Errorf("No access to block list: %s", v)
		} else {
			loaded[idx] = true
			log.Infof("Uploaded %d blocked hosts from: %s", cnt, v)
		}
	}
	return list, loaded
}

func UpdateDb(lists []string, repo repo.ProxyService) (list *BlackList, loaded []bool) {

	list, loaded = updateList(lists)

	err := list.saveToDb(repo)
	if err != nil {
		log.Error(err)
	}

	log.Infof("Loaded %d block hosts from database", list.Count())

	return
}

func LoadFromDb(repo repo.ProxyService) *BlackList {

	list, err := NewBlackList()
	if err != nil {
		log.Errorf("Internal error: %v", err)
		return nil
	}

	cnt := list.loadFromDb(repo)
	log.Infof("Loaded %d block hosts from database", cnt)

	return list
}

func UpdateFile(lists []string) (list *BlackList, loaded []bool) {

	list, loaded = updateList(lists)

	err := list.SaveToFile()
	if err != nil {
		log.Error(err)
	}

	log.Infof("Loaded %d block hosts from file", list.Count())

	return
}

func LoadAdBlock() *BlackList {
	list, err := NewBlackList()
	if err != nil {
		log.Errorf("Internal error: %v", err)
		return nil
	}

	cnt, err := list.loadFromFile()
	if err != nil {
		log.Errorf("load block list from file error: %w", err)
		return list
	}
	log.Debugf("Loaded %d block hosts from file", cnt)

	return list
}

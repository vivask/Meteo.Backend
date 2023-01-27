package kit

import (
	"encoding/json"
	"log"
)

// Default https client
var DefaultClient *Client

func InitClient() {
	c, err := NewClient()
	if err != nil {
		log.Fatalf("Can't create https client, error: %v", err)
	}
	InitLeader()
	DefaultClient = c
}

func PostInt(url string, r interface{}) (body []byte, err error) {
	return DefaultClient.PostInt(url, r)
}

func PostFormInt(url, content string, r interface{}) (body []byte, err error) {
	return DefaultClient.PostFormInt(url, content, r)
}

func PostExt(url string, r interface{}) (body []byte, err error) {
	return DefaultClient.PostExt(url, r)
}

func PutInt(url string, r interface{}) (body []byte, err error) {
	return DefaultClient.PutInt(url, r)
}
func PutExt(url string, r interface{}) (body []byte, err error) {
	return DefaultClient.PutExt(url, r)
}
func PutMain(url string, r interface{}) (body []byte, err error) {
	return DefaultClient.PutMain(url, r)
}
func PutBackup(url string, r interface{}) (body []byte, err error) {
	return DefaultClient.PutBackup(url, r)
}

func GetInt(url string) (body []byte, err error) {
	return DefaultClient.GetInt(url)
}
func GetExt(url string) (body []byte, err error) {
	return DefaultClient.GetExt(url)
}

func GetMain(url string) (body []byte, err error) {
	return DefaultClient.GetMain(url)
}
func GetBackup(url string) (body []byte, err error) {
	return DefaultClient.GetBackup(url)
}

func DeleteInt(url string) (body []byte, err error) {
	return DefaultClient.DeleteInt(url)
}
func DeleteExt(url string) (body []byte, err error) {
	return DefaultClient.DeleteExt(url)
}

func isHealthy(body []byte, err error) bool {
	if err != nil {
		return false
	}

	var state string
	err = json.Unmarshal(body, &state)
	if err != nil {
		return false
	}

	return state == "healthy"
}

func IsMainHealthy(url string) bool {
	body, err := GetMain(url)
	return isHealthy(body, err)
}

func IsBackupHealthy(url string) bool {
	body, err := GetBackup(url)
	return isHealthy(body, err)
}

func IsHealthyInt(url string) bool {
	body, err := GetInt(url)
	return isHealthy(body, err)
}

func IsHealthyExt(url string) bool {
	body, err := GetExt(url)
	return isHealthy(body, err)
}

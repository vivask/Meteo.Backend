package kit

import "encoding/json"

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

package kit

import "encoding/json"

func isEmpty(body []byte, err error) bool {
	if err != nil {
		return true
	}

	var empty bool
	err = json.Unmarshal(body, &empty)
	if err != nil {
		return true
	}

	return empty
}

func IsMainLogEmpty(url string) bool {
	body, err := GetMain(url)
	return isEmpty(body, err)
}

func IsBackupLogEmpty(url string) bool {
	body, err := GetBackup(url)
	return isEmpty(body, err)
}

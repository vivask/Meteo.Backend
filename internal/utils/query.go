package utils

import (
	"fmt"
	"strings"
)

func ParseQuery(query string) (action, tableName string, err error) {
	tableName = "unknow"
	words := strings.Split(query, " ")
	action = words[0]
	switch action {
	case "UPDATE":
		tableName = words[1]
	case "INSERT":
		tableName = words[2]
	case "DELETE":
		tableName = words[2]
	default:
		err = fmt.Errorf("unknown SQL query: %s", words[0])
		return
	}
	return action, strings.Replace(tableName, "\"", "", -1), nil
}

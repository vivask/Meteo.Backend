package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
)

func StringToUint32(s string) (uint32, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}

func StringToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}

func StringToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func RightTrunc(str string, n int) string {
	len := len(str)
	return str[0 : len-n]
}

func LeftTrunc(str string, n int) string {
	len := len(str)
	return str[n : len-1]
}

func GetMD5FileSum(fName string) (string, error) {
	f, err := os.OpenFile(fName, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fSum := md5.New()
	_, err = io.Copy(fSum, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", fSum.Sum(nil)), nil
}

func GetMD5StringlnSum(str string) string {
	b := []byte(str + "\n")
	return fmt.Sprintf("%X", md5.Sum(b))
}

func StringToBool(s string) bool {
	return s == "true"
}

package utils

import (
	"fmt"
	"time"

	"crypto/md5"
	"hash/fnv"
)

func HashString64(str string) uint64 {
	return fnvHashString64(str)
}

func HashString32(str string) uint32 {
	return fnvHashString32(str)
}

func HashTime(dt time.Time) string {
	dts := dt.Format("2006-01-02 15:04:05")
	hasher := md5.New()
	hasher.Write([]byte(dts))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func HashNow32() uint32 {
	dts := time.Now().Format("2006-01-02 15:04:05")
	return fnvHashString32(dts)
}

func fnvHashString32(str string) uint32 {
	h := fnv.New32()
	data := []byte(str)
	_, err := h.Write(data)
	if err != nil {
		return 0
	}
	return h.Sum32()
}

func fnvHashString64(str string) uint64 {
	h := fnv.New64()
	data := []byte(str)
	_, err := h.Write(data)
	if err != nil {
		return 0
	}
	return h.Sum64()
}

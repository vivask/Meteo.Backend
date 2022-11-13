package v1

func toUint8(b bool) (ret uint8) {
	ret = 0
	if b {
		ret = 1
	}
	return
}

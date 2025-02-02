package generate

import (
    "time"
    "math/rand"
    "unsafe"
)

const letterBytes = "abcdefghijhlmnopqrstuvwxyzABCDEFGHIJHLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6
    letterIdxMask = 1<<letterIdxBits - 1
    letterIdxMax = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func NewLinkId() string {
    n := 7
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

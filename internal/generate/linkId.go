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

// NewLinkId generates a 7 letter random string
//
// Implementation from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go#31832326
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

package gails

import (
	"math/rand"
	"time"
)

func RandomStr(sl int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	rst := make([]byte, sl)
	for i := 0; i < sl; i++ {
		rst[i] = chars[rand.Intn(len(chars))]
	}
	return string(rst)
}

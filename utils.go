package gails

import (
	"math/rand"
	"os"
	"os/exec"
	"syscall"
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

//Shell 终端命令操作
func Shell(cmd string, args ...string) error {
	bin, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	return syscall.Exec(bin, append([]string{cmd}, args...), os.Environ())
}

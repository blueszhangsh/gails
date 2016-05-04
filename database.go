package gails

import (
	"strconv"

	"github.com/spf13/viper"
)

func psql(sql ...string) (string, []string) {
	ext := viper.GetStringMap("database.extras")
	args := []string{"-U", ext["user"].(string)}
	if host, ok := ext["host"]; ok {
		args = append(args, "-h", host.(string))
	}
	if port, ok := ext["port"]; ok {
		args = append(args, "-p", strconv.Itoa(port.(int)))
	}
	return "psql", append(args, sql...)
}

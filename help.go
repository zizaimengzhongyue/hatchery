package main

import (
	"flag"
	"fmt"
)

const (
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
)

func SetColor(msg string, color int) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, msg)
}

func init() {
	flag.String("help", "帮助", SetColor("显示帮助信息", Magenta))
	flag.Parse()
}

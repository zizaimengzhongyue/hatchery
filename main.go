package main

import (
	"os"

	"github.com/zizaimengzhongyue/hatchery/logging"
	"github.com/zizaimengzhongyue/hatchery/server"
)

func init() {
	cur, _ := os.Getwd()
	access, err := os.OpenFile(cur+"/log/access.log", os.O_RDONLY|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	logging.New("access", access, "nginx")
	hatchery, err := os.OpenFile(cur+"/log/hatchery.log", os.O_RDONLY|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	logging.New("hatchery", hatchery, "normal")
}

func main() {
	server.Init()
}

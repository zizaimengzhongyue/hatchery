package main

import (
	"net/http"

	"github.com/zizaimengzhongyue/hatchery/server"
)

func main() {
	http.HandleFunc("/register", server.Register)
	http.HandleFunc("/cancel", server.Cancel)
	http.HandleFunc("/get", server.Get)
	http.HandleFunc("/multiGet", server.MultiGet)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}

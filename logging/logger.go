package logging

import (
	"io"
	"log"
	"sync"
)

var logs map[string]*log.Logger = map[string]*log.Logger{}
var mu sync.Mutex = sync.Mutex{}

func New(name string, out io.Writer, format string) {
	mu.Lock()
	defer mu.Unlock()
	logs[name] = log.New(out, "\r", log.Ldate|log.Ltime)
}

func Get(name string) *log.Logger {
	return logs[name]
}

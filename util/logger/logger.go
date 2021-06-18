package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type logger struct {
	Time  time.Time   `json:"time"`
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
	Stack interface{} `json:"stack"`
}

const (
	TypeError = "error"
	TypeInfo  = "info"
	TypeDebug = "debug"
)

func (l *logger) log() {
	fmt.Println(l)
}

func (l *logger) String() string {
	s, _ := json.Marshal(l)
	return string(s)
}

func Error(data interface{}) {
	l := logger{
		Time: time.Now(),
		Type: TypeError,
		Data: data,
	}

	l.log()
}

func Info(data interface{}) {
	l := logger{
		Time: time.Now(),
		Type: TypeInfo,
		Data: data,
	}

	l.log()
}

func Debug(data interface{}) {
	l := logger{
		Time: time.Now(),
		Type: TypeDebug,
		Data: data,
	}

	l.log()
}

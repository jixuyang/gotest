package gotest

import (
	"fmt"
)

const (
	module = "github.com/jixuyang/gotest"
	tag    = "v1.0.1"
)

func GetInfo() string {
	msg := fmt.Sprintf("[%s] [%s] GetInfo", tag, module)
	return msg
}

func GetLicense() string {
	msg := fmt.Sprintf("[%s] [%s] GetLicense", tag, module)
	return msg
}

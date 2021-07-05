package gotest

import (
	"fmt"
)

const (
	module = "gotest"
	tag    = "v0.1.2"
)

func GetInfo() string {
	msg := fmt.Sprintf("[%s] [%s] GetInfo", tag, module)
	return msg
}

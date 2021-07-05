package gotest

import (
	"fmt"
)

const (
	module = "gotest"
)

func GetInfo() string {
	tag := "v0.1.0-alpha"
	msg := fmt.Sprintf("[%s] [%s] GetInfo", tag, module)
	return msg
}

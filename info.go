package gotest

import (
	"fmt"
)

const (
	module = "github.com/jixuyang/gotest"
	tag    = "v1.1.3"
)

func GetInfo() string {
	msg := fmt.Sprintf("[%s] [%s] GetInfo: return module info", tag, module)
	return msg
}

func GetLicense() string {
	msg := fmt.Sprintf("[%s] [%s] GetLicense: return module license", tag, module)
	return msg
}

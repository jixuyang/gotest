package gotest

import (
	"fmt"
)

const (
	module = "github.com/jixuyang/gotest"
	tag    = "v2.0.0"
)

func GetInfo() (string, error) {
	msg := fmt.Sprintf("[%s] [%s] GetInfo: return module info", tag, module)
	return msg, nil
}

func GetLicense() (string, error) {
	msg := fmt.Sprintf("[%s] [%s] GetLicense: return module license", tag, module)
	return msg, nil
}

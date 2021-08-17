package gotest

import (
	"fmt"

	api "github.com/jixuyang/gotest/v2/handler"
	log "github.com/sirupsen/logrus"
)

const (
	module = "github.com/jixuyang/gotest"
	tag    = "v2.1.2"
)

func init() {
	log.Infof("[%s] [%s] init", tag, module)
}

// GetInfo get info
func GetInfo() (string, error) {
	msg := fmt.Sprintf("[%s] [%s] GetInfo: return module info", tag, module)
	return msg, nil
}

// GetLicense get license
func GetLicense() (string, error) {
	msg := fmt.Sprintf("[%s] [%s] GetLicense: return module license", tag, module)
	return msg, nil
}

// GetConfig get config
func GetConfig() (string, error) {
	msg, err := api.GetConfig()
	return msg, err
}

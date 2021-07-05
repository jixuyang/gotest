package api

import "github.com/jixuyang/gotest/v2/util"

func GetConfig() (string, error) {
	path := "/home/jason/config"
	strRet := util.TrimString(path)
	return strRet, nil
}

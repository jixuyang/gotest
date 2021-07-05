package util

import "strings"

func TrimString(str string) string {
	strRet := strings.Trim(str, "\t\n ")
	return strRet
}

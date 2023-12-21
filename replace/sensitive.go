package replace

import (
	"freeTranslate/util"
	"strings"
)

var Sensitive map[string]string

func init() {
	SetSensitive()
}
func GetSensitive(str string) string {
	for k, v := range Sensitive {
		strings.Replace(str, k, v, -1)
	}
	return str
}
func SetSensitive() {
	ss := util.ReadByLine("Sensitive.txt")
	for _, v := range ss {
		Sensitive[strings.Split(v, ":")[0]] = strings.Split(v, ":")[1]
	}
}

package replace

import (
	"freeTranslate/util"
	"log/slog"
	"strings"
)

var Sensitive = map[string]string{}

func GetSensitive(str string) string {
	for k, v := range Sensitive {
		if strings.Contains(str, k) {
			strings.Replace(str, k, v, -1)
			slog.Debug("替换生效", slog.String("key", strings.Split(v, ":")[0]), slog.String("value", strings.Split(v, ":")[1]))
		}
	}
	return str
}
func SetSensitive() {
	ss := util.ReadByLine("Sensitive.txt")
	for _, v := range ss {
		Sensitive[strings.Split(v, ":")[0]] = strings.Split(v, ":")[1]
	}
}

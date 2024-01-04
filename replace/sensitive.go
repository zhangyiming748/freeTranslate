package replace

import (
	"freeTranslate/sql"
	"log/slog"
	"strings"
)

var Sensitive = map[string]string{}

func GetSensitive(str string) string {
	for k, v := range Sensitive {
		if strings.Contains(str, k) {
			strings.Replace(str, k, v, -1)
			slog.Debug("替换生效")
		}
	}
	return str
}
func SetSensitive() {
	m := new(sql.Sensitive)
	ss := m.GetAll()
	for _, s := range ss {
		slog.Info("加载敏感词", slog.String("before", s.Src), slog.String("after", s.Dst))
		Sensitive[s.Src] = s.Dst
	}
}

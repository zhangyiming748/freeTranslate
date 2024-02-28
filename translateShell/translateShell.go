package translateShell

import (
	"freeTranslate/count"
	"freeTranslate/replace"
	"freeTranslate/sql"
	"freeTranslate/util"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	TIMEOUT = 8 //second
)

func Translate(src string) string {
	//trans -brief ja:zh "私の手の動きに合わせて、そう"
	his := new(sql.History)
	defer func() {
		his.SetOne()
	}()
	bing := make(chan string, 1)
	google := make(chan string, 1)
	if runtime.GOOS == "windows" {
		slog.Warn("windows系统需要在Linux子系统中运行")
		os.Exit(-1)
	}
	from := util.GetVal("shell", "from")
	to := util.GetVal("shell", "to")
	proxy := util.GetVal("shell", "proxy")
	language := strings.Join([]string{from, to}, ":")

	go TransByGoogle(proxy, language, src, google)
	go TransByBing(proxy, language, src, bing)

	var dst string
	select {
	case dst = <-bing:
		his.Source = "bing"
	case dst = <-google:
		his.Source = "google"
	case <-time.After(TIMEOUT * time.Second):
		slog.Error("单词翻译出现严重问题")
	}

	dst = replace.ChinesePunctuation(dst)

	his.From = from
	his.To = to
	his.Src = src
	his.Dst = dst

	count.Add("trans")
	return dst
}

package translateShell

import (
	"freeTranslate/count"
	"freeTranslate/replace"
	"freeTranslate/sql"
	"freeTranslate/util"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Translate(src string) string {
	//trans -brief ja:zh "私の手の動きに合わせて、そう"
	if runtime.GOOS == "windows" {
		slog.Warn("windows系统需要在Linux子系统中运行")
		os.Exit(-1)
	}
	google := make(chan string, 1)
	bing := make(chan string, 1)
	from := util.GetVal("shell", "from")
	to := util.GetVal("shell", "to")
	proxy := util.GetVal("shell", "proxy")
	language := strings.Join([]string{from, to}, ":")
	go searchByGoogle(src, language, proxy, google)
	go searchByBing(src, language, bing)
	dst := src
	select {
	case <-time.After(3 * time.Second):
		slog.Warn("Timeout !")
		dst = src
	case dst = <-google:
		slog.Debug("从Google获取翻译", slog.String("原文", src), slog.String("译文", dst))
	case dst = <-bing:
		slog.Debug("从Bing获取翻译", slog.String("原文", src), slog.String("译文", dst))
	}
	his := new(sql.History)
	his.From = from
	his.To = to
	his.Src = src
	his.Dst = dst
	his.Source = "translate"
	his.SetOne()
	count.Add("trans")
	return dst
}

func searchByGoogle(src, language, proxy string, c chan string) {
	cmd := exec.Command("trans", "-brief", "-proxy", proxy, language, src, "-engine", "google")
	output, err := cmd.CombinedOutput()
	dst := string(output)
	if err != nil || output == nil || strings.Contains(dst, "\u001B") || strings.Contains(dst, "Connectiontimedout.RetryingIPv4connection") {
		slog.Warn("查询失败")
	} else {
		dst = replace.ChinesePunctuation(dst)
		c <- dst
	}
}

func searchByBing(src, language string, c chan string) {
	cmd := exec.Command("trans", "-brief", language, src, "-engine", "bing")
	output, err := cmd.CombinedOutput()
	dst := string(output)
	if err != nil || output == nil || strings.Contains(dst, "\u001B") || strings.Contains(dst, "Connectiontimedout.RetryingIPv4connection") {
		slog.Warn("查询失败")
	} else {
		dst = replace.ChinesePunctuation(dst)
		c <- dst
	}
}

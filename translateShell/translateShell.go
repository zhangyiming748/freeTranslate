package translateShell

import (
	"fmt"
	"freeTranslate/baidu"
	"freeTranslate/count"
	"freeTranslate/replace"
	"freeTranslate/sql"
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
	from := os.Getenv("FROM")
	to := os.Getenv("TO")
	proxy := os.Getenv("PROXY")
	language := strings.Join([]string{from, to}, ":")
	cmd := exec.Command("trans", "-brief", "-proxy", proxy, language, src)
	fmt.Println(cmd)
	output, err := cmd.CombinedOutput()
	dst := string(output)
	if err != nil || output == nil || strings.Contains(string(output), "\u001B") || strings.Contains(string(output), "Connectiontimedout.RetryingIPv4connection") {
		time.Sleep(1 * time.Second)
		slog.Warn("临时使用百度翻译")
		dst = baidu.AskBaidu(src)
		count.Add("baidu")
		return dst
	}
	//dst := string(output)
	dst = replace.ChinesePunctuation(dst)
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

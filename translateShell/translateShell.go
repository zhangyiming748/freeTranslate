package translateShell

import (
	"fmt"
	"freeTranslate/count"
	"freeTranslate/replace"
	"freeTranslate/sql"
	"freeTranslate/util"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	TIMEOUT = 1 //second
)

func Translate(src string) string {
	//trans -brief ja:zh "私の手の動きに合わせて、そう"
	his := new(sql.History)
	defer func() {
		his.SetOne()
	}()
	//bing := make(chan string, 1)
	//google := make(chan string, 1)
	if runtime.GOOS == "windows" {
		slog.Warn("windows系统需要在Linux子系统中运行")
		os.Exit(-1)
	}
	from := util.GetVal("shell", "from")
	to := util.GetVal("shell", "to")
	proxy := util.GetVal("shell", "proxy")
	language := strings.Join([]string{from, to}, ":")

	cmd := exec.Command("trans", "-brief", "-proxy", proxy, language, src)
	slog.Info("Trans", slog.String("命令原文", fmt.Sprintf("%s", cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("命令执行出错", slog.String("错误原文", err.Error()))
	}
	var dst string

	dst = string(output)
	dst = replace.ChinesePunctuation(dst)

	his.From = from
	his.To = to
	his.Src = src
	his.Dst = dst

	count.Add("trans")
	return dst
}

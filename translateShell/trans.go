package translateShell

import (
	"fmt"
	"freeTranslate/baidu"
	"freeTranslate/model"
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
	from := util.GetVal("shell", "from")
	to := util.GetVal("shell", "to")
	proxy := util.GetVal("shell", "proxy")
	language := strings.Join([]string{from, to}, ":")
	cmd := exec.Command("trans", "-brief", "-proxy", proxy, language, src)
	fmt.Println(cmd)
	output, err := cmd.CombinedOutput()
	dst := string(output)
	if err != nil || output == nil || strings.Contains(string(output), "\u001B") || strings.Contains(string(output), "Connectiontimedout.RetryingIPv4connection") {
		time.Sleep(1 * time.Second)
		slog.Warn("临时使用百度翻译")
		dst = baidu.AskBaidu(src)
		return dst
	}
	//dst := string(output)
	dst = strings.Replace(dst, "\n", "", 1)
	slog.Debug("翻译成功", slog.String("原文", src), slog.String("译文", dst))
	his := new(model.History)
	his.From = util.GetVal("mysql", "from")
	his.To = util.GetVal("mysql", "to")
	his.Src = src
	his.Dst = dst
	his.Source = "translate-shell"
	if one, err := his.InsertOne(); err != nil {
		slog.Warn("数据库插入新条目失败", slog.String("源", src), slog.String("目标", dst), slog.Any("错误原文", err))
	} else {
		slog.Debug("成功插入缓存到数据库", slog.Int64("条目", one))
	}
	return dst
}

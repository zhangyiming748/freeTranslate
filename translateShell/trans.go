package translateShell

import (
	"fmt"
	"freeTranslate/baidu"
	"freeTranslate/replace"
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
	dst = replace.ChinesePunctuation(dst)
	return dst
}

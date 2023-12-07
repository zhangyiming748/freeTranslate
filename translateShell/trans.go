package translateShell

import (
	"fmt"
	"freeTranslate/util"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	if err != nil {
		slog.Warn("翻译发生错误,是否安装了translate-shell?", slog.Any("错误原文", err))
		os.Exit(-1)
	}
	dst := string(output)
	dst = strings.Replace(dst, "\n", "", 1)
	slog.Debug("翻译成功", slog.String("原文", src), slog.String("译文", dst))
	return dst
}

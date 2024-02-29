package translateShell

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"time"
)

func TransByBing(proxy, language, src string, ans chan string) {
	cmd := exec.Command("trans", "-brief", "-proxy", proxy, language, src)
	slog.Debug("Bing", slog.String("命令原文", fmt.Sprintf("%s", cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		slog.Error("命令执行出错", slog.String("错误原文", err.Error()))
		time.Sleep(3 * time.Second)
		TransByBing(proxy, language, src, ans)
	}
	ans <- string(output)
	close(ans)
}

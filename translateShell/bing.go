package translateShell

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
)

func TransByBing(proxy, language, src string, ans chan string) {
	cmd := exec.Command("trans", "-brief", "-proxy", proxy, language, src)
	slog.Info("Bing", slog.String("命令原文", fmt.Sprintf("%s", cmd)))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	} else if err != nil || output == nil || strings.Contains(string(output), "\u001B") || strings.Contains(string(output), "Connectiontimedout.RetryingIPv4connection") {
		ans <- string(output)
	}
}

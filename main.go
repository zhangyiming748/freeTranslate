package main

import (
	"fmt"
	sql "freeTranslate/sql"
	"freeTranslate/translateShell"
	"freeTranslate/util"
	"io"
	"log/slog"
	"os"
	"time"
)

func init() {
	setLog()
	sql.SetEngine()
}

var fresh []string

func main() {
	//cache := make(map[string]string)
	outname := "after.srt"
	before := util.ReadByLine("before.srt")
	after, _ := os.OpenFile(outname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		after.WriteString(fmt.Sprintf("%s\n", before[i]))
		after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]
		var dst string
		cache := new(sql.History)
		cache.Src = src
		if result := cache.FindOneBySrc(); result.Error == nil {
			dst = cache.Dst
			slog.Debug("find in cache")
		} else {
			dst = translateShell.Translate(src)
			time.Sleep(1 * time.Second)
		}
		after.WriteString(fmt.Sprintf("%s\n", dst))
		fresh = append(fresh, dst)

		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
	}
	os.Remove("before.srt")
	os.Rename("after.srt", "before.srt")

}

func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	file := "Process.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

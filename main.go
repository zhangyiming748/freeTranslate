package main

import (
	"fmt"
	"freeTranslate/baidu"
	"freeTranslate/model"
	"freeTranslate/replace"
	"freeTranslate/storage/mysql"
	"freeTranslate/translateShell"
	"freeTranslate/util"
	"io"
	"log/slog"
	"os"
	"time"
)

func init() {
	setLog()
	//初始化数据库和数据表
	mysql.SetEngine()
	model.SyncHistory()
}

var fresh []string

func main() {
	var (
		c  = 0
		nc = 0
	)
	runLevel := util.GetVal("main", "level")
	//cache := make(map[string]string)
	outname := "after.srt"
	before := util.ReadByLine("before.srt")
	after, _ := os.OpenFile(outname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		after.WriteString(fmt.Sprintf("%s\n", before[i]))
		after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]
		cache := new(model.History)
		cache.Src = src
		if has, _ := cache.FindBySrc(); has {
			after.WriteString(fmt.Sprintf("%s\n", cache.Dst))
			fmt.Println("找到缓存,直接返回")
			c++
		} else {
			var dst string
			switch runLevel {
			case "baidu":
				dst = baidu.AskBaidu(src)
			case "trans":
				dst = translateShell.Translate(src)
			}
			dst = replace.ChinesePunctuation(dst)
			after.WriteString(fmt.Sprintf("%s\n", dst))

			nc++
			fresh = append(fresh, dst)
			time.Sleep(2 * time.Second)
		}
		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
	}
	slog.Info("翻译结束", slog.Int("从缓存中找到", c), slog.Int("新建查询", nc), slog.Any("其中增加的单词", fresh))
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

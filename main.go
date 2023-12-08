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
func main() {
	var (
		c  = 0
		nc = 0
	)
	runLevel := util.GetVal("main", "level")
	//cache := make(map[string]string)
	outname := "after.srt"
	switch runLevel {
	case "baidu":

		outname = "baidu.srt"
	case "trans":
		outname = "trans.srt"
	}

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
			his := new(model.History)
			his.From = util.GetVal("mysql", "from")
			his.To = util.GetVal("mysql", "to")
			his.Src = src
			his.Dst = dst
			if one, err := his.InsertOne(); err != nil {
				slog.Warn("数据库插入新条目失败", slog.String("源", src), slog.String("目标", dst), slog.Any("错误原文", err))
			} else {
				slog.Debug("成功插入缓存到数据库", slog.Int64("条目", one))
			}
			nc++
			time.Sleep(2 * time.Second)
		}
		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
	}
	slog.Info("翻译结束", slog.Int("从缓存中找到", c), slog.Int("新建查询", nc))
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

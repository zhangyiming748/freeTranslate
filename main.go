package main

import (
	"fmt"
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
	//cache := make(map[string]string)
	before := util.ReadByLine("before.srt")
	after, _ := os.OpenFile("after.srt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		after.WriteString(fmt.Sprintf("%s\n", before[i]))
		after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]
		cache := new(model.History)
		cache.Src = src
		if has, _ := cache.FindBySrc(); has {
			after.WriteString(fmt.Sprintf("%s\n", cache.Dst))
		} else {
			//dst := baidu.AskBaidu(src)
			dst := translateShell.Translate(src)
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
		}
		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
		time.Sleep(10 * time.Second)
		for t := 2; t > 0; t-- {
			fmt.Printf("冷却时间还有%d秒\n", t)
			time.Sleep(time.Second)
		}
		//fmt.Printf("循环一次后cache的情况: %+v\n", cache)
	}
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

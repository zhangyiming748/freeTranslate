package main

import (
	"fmt"
	"freeTranslate/baidu"
	"freeTranslate/model"
	"freeTranslate/storage/mysql"
	"freeTranslate/util"
	"log/slog"
	"os"
	"time"
)

func init() {
	//初始化数据库和数据表
	mysql.SetEngine()
	model.SyncHistory()
}
func main() {
	//cache := make(map[string]string)
	before := util.ReadByLine("before.srt")
	after, _ := os.OpenFile("after.srt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		after.WriteString(fmt.Sprintf("序号 : %s\n", before[i]))
		after.WriteString(fmt.Sprintf("时间轴 : %s\n", before[i+1]))
		src := before[i+2]
		cache := new(model.History)
		cache.Src = src
		if has, _ := cache.FindBySrc(); has {
			after.WriteString(fmt.Sprintf("字幕 : %s\n", cache.Dst))
		} else {
			dst := baidu.AskBaidu(src)
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

		//if dst, ok := cache[src]; ok {
		//	after.WriteString(fmt.Sprintf("字幕 : %s\n", dst))
		//} else {
		//	dst = baidu.AskBaidu(src)
		//	//dst = "c"
		//	cache[src] = dst
		//	after.WriteString(fmt.Sprintf("字幕 : %s\n", dst))
		//	h := new(model.History)
		//	h.Src = src
		//	h.Dst = dst
		//	_, err := h.InsertOne()
		//	if err != nil {
		//		fmt.Println("数据库插入错误")
		//		continue
		//	}
		//}
		after.WriteString(fmt.Sprintf("空行 : %s\n", before[i+3]))
		after.Sync()
		time.Sleep(10 * time.Second)
		//fmt.Printf("循环一次后cache的情况: %+v\n", cache)
	}
}

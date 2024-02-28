package main

import (
	"fmt"
	"freeTranslate/GetAllFolder"
	"freeTranslate/GetFileInfo"
	"freeTranslate/count"
	"freeTranslate/replace"
	sql "freeTranslate/sql"
	"freeTranslate/translateShell"
	"freeTranslate/util"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {
	setLog()
	sql.SetEngine()
}

var fresh []string

func main() {
	replace.SetSensitive()
	folders := GetAllFolder.List(util.GetVal("root", "dir"))
	folders = append(folders, util.GetVal("root", "dir"))
	for _, folder := range folders {
		files := GetFileInfo.GetAllFileInfo(folder, "srt")
		for _, file := range files {
			if strings.Contains(file.PurgeName, "origin") {
				continue
			}
			trans(file.FullPath)
		}
	}
	cpdatabase()

}
func trans(srt string) {
	//host := strings.Split(util.GetVal("shell", "proxy"), ":")[0]
	//port := strings.Split(util.GetVal("shell", "proxy"), ":")[1]
	host, port, _ := net.SplitHostPort(util.GetVal("shell", "proxy"))
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second*5); err != nil {
		slog.Error("代理网址不可达")
		os.Exit(-1)
	} else {
		slog.Info("", slog.Any("本地连接", conn.LocalAddr()), slog.Any("远程连接", conn.RemoteAddr()))
	}

	seed := rand.New(rand.NewSource(time.Now().Unix()))
	r := seed.Intn(2000)
	//中间文件名
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadByLine(srt)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		after.WriteString(fmt.Sprintf("%s\n", before[i]))
		after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]

		afterSrc := replace.GetSensitive(src)

		var dst string
		cache := new(sql.History)
		cache.Src = afterSrc
		if result := cache.FindOneBySrc(); result.Error == nil {
			dst = cache.Dst
			slog.Debug("find in cache")
			count.Add("cache")
		} else {
			dst = translateShell.Translate(afterSrc)
			time.Sleep(1 * time.Second)
		}
		dst = replace.GetSensitive(dst)
		slog.Info("", slog.String("文件名", tmpname), slog.String("原文", src), slog.String("译文", dst))
		after.WriteString(fmt.Sprintf("%s\n", src))
		after.WriteString(fmt.Sprintf("%s\n", dst))
		fresh = append(fresh, dst)

		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
	}
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	exec.Command("cp", srt, origin).CombinedOutput()
	os.Rename(tmpname, srt)
	t, b, c := count.Get()
	slog.Info("统计", slog.Int("缓存", c), slog.Int("百度", b), slog.Int("shell", t))

}
func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	//file := "trans.log"
	//logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	//if err != nil {
	//	panic(err)
	//}
	//logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opt))
	slog.SetDefault(logger)
}
func cpdatabase() {
	folderPath := "/data"
	_, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在")
	} else if err != nil {
		fmt.Println("发生错误：", err)
	} else {
		fmt.Println("文件夹存在")
		exec.Command("cp", "trans", "/data").Run()
	}
}

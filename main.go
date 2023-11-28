package main

import (
	"fmt"
	"freeTranslate/util"
	"os"
)

var Cache map[string]string

func main() {
	before := util.ReadByLine("before.srt")
	after, _ := os.OpenFile("after.srt", 2|8|512, 0777)
	for i := 0; i < len(before); i += 4 {
		after.WriteString(fmt.Sprintf("序号 : %s\n", before[i]))

		after.WriteString(fmt.Sprintf("时间轴 : %s\n", before[i+1]))

		after.WriteString(fmt.Sprintf("字幕 : %s\n", before[i+2]))

		after.WriteString(fmt.Sprintf("空行 : %s\n", before[i+3]))
	}
	//baidu.AskBaidu("", "", "")
}

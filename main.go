package main

import (
	"fmt"
	"freeTranslate/util"
)

func main() {
	before := util.ReadByLine("before.srt")
	for _, line := range before {
		fmt.Println(line)
	}
	//baidu.AskBaidu("", "", "")
}

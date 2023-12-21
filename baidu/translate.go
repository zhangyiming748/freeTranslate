package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"freeTranslate/constant"
	"freeTranslate/replace"
	"freeTranslate/sql"
	"freeTranslate/util"
	"log/slog"
	"os"
	"strings"
)

type Success struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}
type Failure struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func AskBaidu(query string) string {
	appid := util.GetVal("baidu", "appid")
	if eappid := os.Getenv("APPID"); eappid != "" {
		appid = eappid
		slog.Info("使用环境变量的appid")
	}
	key := util.GetVal("baidu", "key")
	if ekey := os.Getenv("APPID"); ekey != "" {
		key = ekey
		slog.Info("使用环境变量的key")
	}
	from := constant.T2B[util.GetVal("shell", "from")]
	to := constant.T2B[util.GetVal("shell", "to")]

	salt := GetSalt()
	after := strings.Join([]string{appid, query, salt, key}, "") //拼接字符串
	sign := GetMD5(after)                                        // 计算签名
	//sign = url.QueryEscape(sign)
	param := make(map[string]string)
	param["q"] = query
	param["from"] = from
	param["to"] = to
	param["appid"] = appid
	param["salt"] = salt
	param["sign"] = sign
	fmt.Printf("param : %+v\n", param)
	get, err := util.HttpGet(nil, param, constant.HTTPS)
	if err != nil {
		return ""
	}
	var s Success
	var f Failure
	if err = json.Unmarshal(get, &s); err != nil {
		err = json.Unmarshal(get, &f)
		if err != nil {
			slog.Warn("两种结构体序列化均失败")
			return ""
		}
		slog.Debug("失败", slog.Any("结构体", f))
		h := new(sql.History)
		h.ErrorCode = f.ErrorCode
		h.ErrorMsg = f.ErrorMsg
		h.Src = query
		h.From = from
		h.To = to
		h.SetOne()
	}
	dst := replace.ChinesePunctuation(s.TransResult[0].Dst)
	dst = replace.ChinesePunctuation(dst)
	his := new(sql.History)
	his.From = from
	his.To = to
	his.Src = query
	his.Dst = dst
	his.Source = "baidu"
	var parameters string
	for k, v := range param {
		parameters = strings.Join([]string{k, v}, ":")
	}
	his.Request = parameters
	his.SetOne()
	slog.Debug("翻译成功", slog.String("原文", query), slog.String("译文", dst))
	return dst
}

/*
计算MD5
*/
func GetMD5(str string) string {
	hash := md5.Sum([]byte(str))
	md5Str := hex.EncodeToString(hash[:])
	return md5Str
}

package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"freeTranslate/constant"
	"freeTranslate/model"
	"freeTranslate/replace"
	"freeTranslate/util"
	"log/slog"
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
	key := util.GetVal("baidu", "key")
	from := util.GetVal("baidu", "from")
	to := util.GetVal("baidu", "to")

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
		h := new(model.History)
		h.ErrorCode = f.ErrorCode
		h.ErrorMsg = f.ErrorMsg
		h.Src = query
		h.From = from
		_, err = h.InsertOne()
		if err != nil {
			slog.Warn("数据库插入新条目失败", slog.String("源", query), slog.Any("错误原文", err))
			return ""
		}
	} else {
		slog.Debug("成功", slog.Any("结构体", s))
	}
	resule := replace.ChinesePunctuation(s.TransResult[0].Dst)
	slog.Debug("翻译成功", slog.String("原文", query), slog.String("译文", resule))
	return resule
}

/*
计算MD5
*/
func GetMD5(str string) string {
	hash := md5.Sum([]byte(str))
	md5Str := hex.EncodeToString(hash[:])
	return md5Str
}

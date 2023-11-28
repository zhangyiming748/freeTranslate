package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"freeTranslate/constant"
	"freeTranslate/replace"
	"freeTranslate/util"
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
	err = json.Unmarshal(get, &s)
	fmt.Printf("得到的结构体%+v\n", s)
	fmt.Println(string(get))
	resule := replace.ChinesePunctuation(s.TransResult[0].Dst)
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

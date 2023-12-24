package constant

var (
	BitRate = map[string]string{
		"avc":  "5000K",
		"hevc": "1800K",
	}
)

const (
	Type      = iota + 1
	Kilobyte  = 1000 * Type
	Megabyte  = 1000 * Kilobyte
	Gigabyte  = 1000 * Megabyte
	Terabyte  = 1000 * Gigabyte
	Petabyte  = 1000 * Terabyte
	Exabyte   = 1000 * Petabyte
	Zettabyte = 1000 * Exabyte
	Yottabyte = 1000 * Zettabyte
)
const HTTPS = "https://fanyi-api.baidu.com/api/trans/vip/translate"
const (
	JP = "jp"
	EN = "en"
	ZH = "zh"
)

var T2B = map[string]string{
	"en":    "en",
	"ja":    "jp",
	"zh-CN": "zh",
	"ko":    "kor", // 韩语
	"th":    "th",  // 泰语
	"de":    "de",  //德语
	"fr":    "fra", //法语
	"ru":    "ru",  // 俄语
	"sp":    "spa", // 西班牙语
}

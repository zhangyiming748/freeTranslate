package count

var (
	trans int
	baidu int
	cache int
)

func Add(from string) {
	switch from {
	case "trans":
		trans++
	case "baidu":
		baidu++
	case "cache":
		cache++
	}
}
func Get() (trans, baidu, cache int) {
	return trans, baidu, cache
}

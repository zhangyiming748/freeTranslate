package count

var (
	trans int
	cache int
)

func Add(from string) {
	switch from {
	case "trans":
		trans++
	case "cache":
		cache++
	}
}
func Get() (trans, cache int) {
	return trans, cache
}

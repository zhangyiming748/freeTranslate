package baidu

import (
	"testing"
)

func TestJp(t *testing.T) {
	ret := AskBaidu("jp", "zh", "あっ、すいません")
	t.Log(ret)
}
func TestEn(t *testing.T) {
	AskBaidu("zh", "en", "苹果")
}
func TestStep1(t *testing.T) {
	AskBaidu("zh", "en", "苹果")
}

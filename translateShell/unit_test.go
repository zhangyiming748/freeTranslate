package translateShell

import (
	"freeTranslate/replace"
	"testing"
)

// go test -v -run TestTranslate
func TestTranslate(t *testing.T) {
	dst := Translate("私の手の動きに合わせて、そう")
	dst = replace.ChinesePunctuation(dst)
	t.Log(dst)
}

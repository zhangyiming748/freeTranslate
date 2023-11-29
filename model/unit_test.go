package model

import (
	"freeTranslate/storage/mysql"
	"testing"
)

func init() {
	mysql.SetEngine()
	SyncHistory()
}

func TestFind(t *testing.T) {
	h := new(History)
	h.Src = "あー"
	if has, _ := h.FindBySrc(); has {
		t.Log("find!")
	} else {
		t.Log("not!")
	}
	t.Logf("%+v\n", h)
}
func TestFindNot(t *testing.T) {
	h := new(History)
	h.Src = "3233"
	if has, _ := h.FindBySrc(); has {
		t.Log("find!")
	} else {
		t.Log("not!")
	}
	t.Logf("%+v\n", h)
}

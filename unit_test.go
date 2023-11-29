package main

import (
	"freeTranslate/model"
	"freeTranslate/storage/mysql"
	"log/slog"
	"testing"
)

func init() {
	mysql.SetEngine()
	model.SyncHistory()
}
func TestUnit(t *testing.T) {
	mysql.SetEngine()
	h := new(model.History)
	all, err := h.Sum()
	if err != nil {
		return
	}
	slog.Info("all words", slog.Int64("共翻译单词", all))
}

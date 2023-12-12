package sql

import (
	"fmt"
	"testing"
)

func TestSetOne(t *testing.T) {
	SetEngine()
	h := new(History)
	h.From = "en"
	h.To = "zh"
	h.Src = "hello"
	h.Dst = "你好"
	h.Source = "myself"
	q := h.SetOne()
	t.Logf("%+v\n", q)
}
func TestGetOne(t *testing.T) {
	SetEngine()
	y := new(History)
	y.Src = "hello"
	r1 := y.FindOneBySrc("hello")
	t.Logf("%+v\n", y)
	n := new(History)
	n.Src = "he"
	r2 := n.FindOneBySrc("hello")
	t.Logf("%+v\n", n)
	fmt.Printf("r1 = %v\nr2 = %v\n", r1, r2)
	fmt.Printf("error = %v\nrowsAffected = %v\nstatement = %v\nclone = %v\n", r2.Error, r2.RowsAffected, r2.Statement, r2.Row())
	//error = record not found
}

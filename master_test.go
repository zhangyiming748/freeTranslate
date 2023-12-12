package main

import (
	"fmt"
	"freeTranslate/sql"
	"testing"
)

func TestSetOne(t *testing.T) {
	sql.SetEngine()
	h := new(sql.History)
	h.From = "en"
	h.To = "zh"
	h.Src = "hello"
	h.Dst = "你好"
	h.Source = "myself"
	q := h.SetOne()
	t.Logf("%+v\n", q)
}
func TestGetOne(t *testing.T) {
	sql.SetEngine()
	y := new(sql.History)
	y.Src = "hello"
	r1 := y.FindOneBySrc()
	t.Logf("%+v\n", y)
	n := new(sql.History)
	n.Src = "hello"
	r2 := n.FindOneBySrc()
	fmt.Println(r2)
	t.Logf("%+v\n")
	fmt.Printf("r1 = %v\nr2 = %v\n", r1, r2)

	fmt.Printf("error = %v\nrowsAffected = %v\nstatement = %v\nclone = %v\n", r2.Error, r2.RowsAffected, r2.Statement.Context, r2.Row())
	//error = record not found
}

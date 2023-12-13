package main

import (
	"encoding/json"
	"fmt"
	"freeTranslate/sql"
	"os"
	"testing"
)

type S []Mysql
type Mysql struct {
	Id         int         `json:"id"`
	From       string      `json:"from"`
	To         string      `json:"to"`
	Src        string      `json:"src"`
	Dst        string      `json:"dst"`
	ErrorCode  string      `json:"error_code"`
	ErrorMsg   string      `json:"error_msg"`
	Request    string      `json:"request"`
	UpdateTime string      `json:"update_time"`
	CreateTime string      `json:"create_time"`
	DeleteTime interface{} `json:"delete_time"`
	Source     *string     `json:"source"`
}

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
	//sql.SetEngine()
	y := new(sql.History)
	y.Src = "hello"
	r1 := y.FindOneBySrc()
	t.Logf("%+v\n", y)
	n := new(sql.History)
	n.Src = "hello"
	r2 := n.FindOneBySrc()
	fmt.Println(r2)
	t.Logf("%+v\n", r2)
	fmt.Printf("r1 = %v\nr2 = %v\n", r1, r2)

	fmt.Printf("error = %v\nrowsAffected = %v\nstatement = %v\nclone = %v\n", r2.Error, r2.RowsAffected, r2.Statement.Context, r2.Row())
	//error = record not found
}

// go test -v -run MysqlJson2Sql
func TestMysqlJson2Sql(t *testing.T) {
	file, err := os.ReadFile("/home/zen/git/freeTranslate/history.json")
	if err != nil {
		return
	}

	var s S
	json.Unmarshal(file, &s)
	//t.Log(s)

	for _, row := range s {
		fmt.Printf("from : %v\nto : %v\nsrc : %v\ndet : %v\n", row.From, row.To, row.Src, row.Dst)
		var h sql.History
		h.Src = row.Src
		h.Dst = row.Dst
		h.SetOne()
	}

}

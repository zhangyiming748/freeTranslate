package sql

import (
	"testing"
)

// go test -v -run TestCreate

func TestCreate(t *testing.T) {
	SetEngine()
	s := new(Sensitive)
	ret := s.GetAll()
	t.Logf("%+v\n", ret)
	for _, v := range ret {
		t.Logf("src = %s\tdst = %s\n", v.Src, v.Dst)

	}
}

func TestAddOne(t *testing.T) {
	SetEngine()
	s := new(Sensitive)
	s.Src = "1"
	s.Dst = "2"
	s.AddOne()
}

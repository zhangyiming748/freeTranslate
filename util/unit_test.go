package util

import (
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

var (
	concurrent      int32
	concurrentLimit = make(chan struct{}, 10)
)

func ReadDB() {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("readDB并发度%v\n", atomic.LoadInt32(&concurrent))
	time.Sleep(200 * time.Microsecond)
	atomic.AddInt32(&concurrent, -1)
}
func handler(f func()) {
	concurrentLimit <- struct{}{}
	//readDB()
	f()
	<-concurrentLimit
	return
}
func TestLimitThreads(t *testing.T) {
	for i := 0; i < 100; i++ {
		go handler(ReadDB)
	}
	time.Sleep(3 * time.Second)
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomWithSeed()
	}
}
func TestDuplicateBySlice(t *testing.T) {
	s := ReadByLine("E:\\git\\ProcessAVI\\util\\list.txt")
	r := DuplicateBySlice(s)
	result := []string{}
	for _, v := range r {
		prefix := "ytdlp --proxy 127.0.0.1:8889"
		suffix := strings.Join([]string{prefix, v}, " ")
		result = append(result, suffix)
	}
	WriteByLine("E:\\git\\ProcessAVI\\util\\plist.ps1", result)
}

func TestSave(t *testing.T) {
	size, err := GetSize("/Users/zen/Downloads/curl.go")
	if err != nil {
		return
	}
	t.Log(size)
}

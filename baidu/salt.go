package baidu

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
生成随机数
*/
var seed = rand.New(rand.NewSource(time.Now().Unix()))

func GetSalt() string {
	a := strconv.Itoa(seed.Intn(10))
	b := strconv.Itoa(seed.Intn(10))
	c := strconv.Itoa(seed.Intn(10))
	d := strconv.Itoa(seed.Intn(10))
	e := strconv.Itoa(seed.Intn(10))
	f := strconv.Itoa(seed.Intn(10))
	g := strconv.Itoa(seed.Intn(10))
	h := strconv.Itoa(seed.Intn(10))
	i := strconv.Itoa(seed.Intn(10))
	j := strconv.Itoa(seed.Intn(10))
	return strings.Join([]string{a, b, c, d, e, f, g, h, i, j}, "")
}

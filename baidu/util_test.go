package baidu

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "e"
	m["b"] = "f"
	if v, ok := m["a"]; ok {
		fmt.Println("have", v)
	} else {
		fmt.Println("have not")
	}
}

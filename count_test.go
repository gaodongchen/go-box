package gobox

import "testing"

// 测试位运算
func Test1Mib(t *testing.T) {
	i := 1 << 20
	if i != 1024*1024 {
		t.Fatal(i)
	}
}

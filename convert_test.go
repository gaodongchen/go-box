package gobox

import "testing"

func TestCBin2deci(t *testing.T) {
	r, err := Bin2deci("111")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(r)
}

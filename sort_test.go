package gobox

import (
	"testing"
)

func TestBubbleSortAsc(t *testing.T) {
	r := BubbleSort([]interface{}{3, 8, 1, 6, 5, 4, 9, 2, 7}, func(i1 interface{}, i2 interface{}) bool {
		return i1.(int) > i2.(int)
	})
	t.Log(r)
}

func TestBubbleSortDesc(t *testing.T) {
	r := BubbleSort([]interface{}{4, 8, 1, 6, 5, 9, 3, 2, 7}, func(i1 interface{}, i2 interface{}) bool {
		return i1.(int) < i2.(int)
	})
	t.Log(r)
}

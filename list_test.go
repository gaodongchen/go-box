package gobox

import "testing"

func TestChunkList(t *testing.T) {
	r := Chunk([]interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, 11)
	t.Log(r)
}

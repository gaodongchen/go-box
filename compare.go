package gobox

// 取比较小的数
func Min(max int, min int) int {
	if min > max {
		return max
	}
	return min
}

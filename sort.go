package gobox

// 冒泡排序
func BubbleSort(l []interface{}, f func(interface{}, interface{}) bool) []interface{} {
	llen := len(l)
	for i := 0; i < llen; i++ {
		for j := llen - 1; j > i; j-- {
			if f(l[j-1], l[j]) {
				lt := l[j-1]
				l[j-1] = l[j]
				l[j] = lt
			}
		}
	}
	return l
}

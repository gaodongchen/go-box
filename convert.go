package gobox

import (
	"strconv"
)

// 字符串二进制变成十进制数字
func Bin2deci(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 2, 64)

	if nil != err {
		return 0, err
	}

	return i, nil
}

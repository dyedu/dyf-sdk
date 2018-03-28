package tools

import (
	"math"
	"regexp"
	"strings"
)

/*
@arg:
	str			需要转移的字符串
@desc:
    通过关键字模糊匹配时用，会转移字符串，并把空格置换成.*
@ret:
    string		转义后字符串
@author:
    caowh create on 2017-3-13
*/
func ParseRegex(str string) string {
	var data = regexp.QuoteMeta(str)
	var array = strings.Split(data, " ")
	var result = ".*"
	for _, item := range array {
		result += item + ".*"
	}
	return result
}

/*
@arg
	f:	目标值
	n:	保留几位小数
@desc
	对目标值保留n位小数
@ret
	float64
@author
	modify by caowh on 2017-06-13	加注释
*/
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

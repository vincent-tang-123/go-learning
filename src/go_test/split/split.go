package split

import (
	"strings"
	"testing"
)

// split.go

// Split 将 s 按照sep进行分割 ，返回一个字符串切片
// Split("我爱你", "爱) => ["我", "你"]
func Split(s, sep string)(ret []string){
	idx := strings.Index(s, sep) // 获取 s 的索引
	for idx > -1{ // 按索引取值，知道索引不满足条件
		ret = append(ret, s[:idx]) // 切片添加赋值给 ret
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}


func BenchmarkSplit(b *testing.B){
	// b.N 不是固定的数
	for i:=0;i<b.N;i++{
		Split("沙河有沙又有河", "沙")
	}
}
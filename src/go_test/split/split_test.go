package split

import (
	"reflect"
	"testing"
)

// 测试

//func TestSplit(t *testing.T) {
//	got := Split("我爱你", "爱")
//	want := []string{"我", "你"}
//
//	//got := Split("a:b:c", ":") // "a:b:c" 按照冒号去分割
//	//want := []string{"a", "b", "c"}
//	if !reflect.DeepEqual(got, want){
//		t.Errorf("want:%v got:%v", want, got)
//	}
//}


func TestSplit(t *testing.T) {
	type test1 struct {
		input string
		sep string
		want []string
	}
	tests := map[string]test1{
		"simple": test1{input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"multi sep": test1{input: "a:b:c", sep:":", want:[]string{"a", "b", "c"}},
		"multi sep2": test1{input:"abcd", sep: "bc", want:[]string{"a", "d"}},
	}
	for name, tc := range tests{
		t.Run(name, func(t *testing.T) { // t.Run 拿到子单元测试详细信息
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want){
				t.Errorf("name: %s falied, want: %v got:%v", name, tc.want, got)
			}
		})
	}
}

/*
终端输入 go test -v

=== RUN   TestSplit
=== RUN   TestSplit/multi_sep2
=== RUN   TestSplit/simple
=== RUN   TestSplit/multi_sep
--- PASS: TestSplit (0.00s)
    --- PASS: TestSplit/multi_sep2 (0.00s)
    --- PASS: TestSplit/simple (0.00s)
    --- PASS: TestSplit/multi_sep (0.00s)
PASS
ok      vincent441/src/go_test/split    0.395s
*/


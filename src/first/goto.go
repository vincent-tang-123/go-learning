package main
import "fmt"
// goto C 语言中一样有goto
// goto 跳跃到指定的行，然后向下执行代码
func main() {
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	if name == "wupeiqi" {
		// svip
		goto SVIP
	} else if name == "yuanhao" {
		// vip
		goto VIP
	}
	fmt.Println("预约：")
	VIP:
		fmt.Println("等号")
	SVIP:
		fmt.Println("进入")
}

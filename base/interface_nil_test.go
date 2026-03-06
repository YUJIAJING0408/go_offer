package base

import (
	"fmt"
	"testing"
)

/*
@Date:
@Auth: YUJIAJING
@Desp: 接口零值与nil判断
*/

func TestInterfaceNil(t *testing.T) {
	var i any             // 定义空接口
	i = nil               // 此时空接口i的类型和值都为nil，i==nil恒成立
	fmt.Println(i == nil) // true
	type User struct {
		Name string
	}
	var u *User = nil     // u==nil恒成立
	fmt.Println(u == nil) // ture
	i = u                 // 此时i==nil就不成立了，因为只有值为nil，接口类型是*User
	fmt.Println(i == nil) // false
}

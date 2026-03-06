package base

import (
	"fmt"
	"testing"
)

/*
@Date:
@Auth: YUJIAJING
@Desp: Go语言使用断言时会发生拷贝吗？
*/

type User struct {
	Name string
	Age  int
}

func TestValueAssert(t *testing.T) {
	var u any = User{Name: "gopher"}
	fmt.Printf("User地址: %p\n", &u)
	if ua, ok := u.(User); ok {
		fmt.Printf("User断言后地址: %p\n", &ua)
	}
}

func TestPointerAssert1(t *testing.T) {
	var u any = &User{Name: "gopher"}
	fmt.Printf("*User地址: %p\n", u)
	if ua, ok := u.(*User); ok {
		fmt.Printf("*User断言后地址: %p\n", ua)
	}
}

func TestPointerAssert2(t *testing.T) {
	/*
		[]int引用类型，在断言时断言结果也为引用类型，故持指向同一片地址
	*/
	var u any = make([]int, 5)
	fmt.Printf("int切片地址: %p\n", u)
	if ua, ok := u.([]int); ok {
		fmt.Printf("int切片断言后地址: %p\n", ua)
	}
}

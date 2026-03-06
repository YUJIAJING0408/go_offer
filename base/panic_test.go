package base

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
@Date:
@Auth: YUJIAJING
@Desp: Go中触发异常的场景有哪些？
*/
/*
Go中通过error类型处理应用中业务逻辑错误，但是可能存在非业务逻辑错误等情况，error不会影响程序继续进行
但除开error外还存在系统级的错误（panic，异常）：
（1）数组切片下标越界
（2）空指针解引用
（3）主动调用panic函数
（4）非法的类型断言
（5）数学错误
（6）内存越界或非法操作
（7）运行时错误
（8）使用不安全的库或代码
需要通过defer和recover可以处理panic，defer确保某些操作始终会被执行，而recover可以用来捕获panic并防止程序崩溃
*/

func TestSlice(t *testing.T) {
	var arr = make([]int, 5)
	fmt.Println(arr)      // [0 0 0 0 0]
	fmt.Println(len(arr)) // 5
	// 尝试访问第六个元素
	fmt.Println(arr[5]) // panic: runtime error: index out of range [5] with length 5 [recovered, repanicked]
}

func TestNilPoint(t *testing.T) {
	//var ptr *int
	//fmt.Println(ptr) // nil
	//// 尝试解引用nil指针
	//fmt.Println(*ptr) // panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]
	var m map[string]string
	fmt.Println(m)        // 不报错 map[]
	fmt.Println(m == nil) // true
	// 尝试写入未初始化的nil Map
	m["a"] = "b" // panic: assignment to entry in nil map
}

func TestPanic(t *testing.T) {
	panic("主动触发Panic")
}

func TestAssert(t *testing.T) {
	var a any = "hello world"
	fmt.Println(a) // "hello world"
	// 错误的断言
	num := a.(int) // panic: interface conversion: interface {} is string, not int [recovered, repanicked]
	fmt.Println(num)
}

func TestNumber(t *testing.T) {
	// 错误的数学运算，甚至无法通过编译器，直接在编译时就检查出错误
	// res := 1 / 0 // .\panic_test.go:59:13: invalid operation: division by zero
	// fmt.Println(res)
}

func TestUnsafePointer(t *testing.T) {
	var p unsafe.Pointer
	*(*int)(p) = 11 // panic: runtime error: invalid memory address or nil pointer dereference
}

func Div(a, b int) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("捕获到异常（%v）", r)
		}
	}()
	res = a / b
	return
}

func TestDeferAndRecover(t *testing.T) {
	if res, err := Div(1, 0); err == nil {
		fmt.Printf("计算结果：%v\n", res)
	} else {
		fmt.Printf("计算错误：%v\n", err)
	}
}

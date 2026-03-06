package base

import (
	"fmt"
	"testing"
)

/*
@Date:
@Auth: YUJIAJING
@Desp: defer的变量快照什么时候会失效
*/
/*
defer变量快照指的被defer语句定义时所捕获的的变量状态。某些情况下会出现快照失效，导致程序不按预期运行：
（1）匿名闭包函数：当defer语句中使用匿名函数捕获外部变量，如果变量值在defer语句定义后发生变化，defer执行时会使用变化后的值。
	解决方法是不要在defer后函数使用外部变量，而是改为显示传值。
（2）引用类型：当defer引用持有引用类型的变量是，岁引用本身地址不变，但指向内容可能会发生变化，这也可能导致快照失效。
*/

func TestClosureExpire(t *testing.T) {
	var num = 0
	defer func() {
		fmt.Println("defer：", num) // 1
	}()
	num++
}
func TestClosureFix(t *testing.T) {
	var num = 0
	defer func(n int) {
		fmt.Println("defer：", n) // 0
	}(num)
	num++
}

func TestPointerExpire(t *testing.T) {
	var s = make([]int, 3) // [0,0,0]
	defer func() {
		fmt.Println(s) // [0,10,0]
	}()
	s[1] = 10
}

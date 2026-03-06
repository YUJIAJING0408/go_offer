package base

import (
	"fmt"
	"testing"
)

/*
@Date:
@Auth: YUJIAJING
@Desp: Go中如何实现闭包，闭包主要适用场景
*/

/*
在Go中闭包是一个函数值，允许引用外部作用域中变量。
实现闭包方法非常简单，在A函数内部定义B函数，B函数体中访问A函数中变量。
闭包的作用非常广泛：
（1）伪全局变量
（2）函数工厂根据不同配置参数动态创建函数
（3）装饰器模式通过闭包在不修改原有函数基础上动态添加新功能
（4）回调函数将A函数作为参数传递给B函数，通过闭包捕获上下文并执行函数
（5）并发编程
（6）保存中间态，如下计数器示例
*/

// 闭包实现计数器
func adder() func(int) int {
	// 外层函数，定义变量
	sum := 0
	return func(x int) int {
		// 内层函数使用外部函数变量sum
		sum += x
		return sum
	}
}

func TestAdder(t *testing.T) {
	adder1, adder2 := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(adder1(i), adder2(2*i)) //使用闭包
	}
}

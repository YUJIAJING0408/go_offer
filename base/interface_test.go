package base

import (
	"fmt"
	"testing"
)

/*
@Date:
@Auth: YUJIAJING
@Desp: Go语言的结构类型是怎么实现的？
*/

/*
Go的接口是一种动态类型，其中最重要的就是空接口interface{}（别名any）
Go的接口本质上是一个动态类型和动态值的组合，这部分需要看接口底层实现。
Go接口采用鸭子类型设计，无需显示声明实现关系，只要一个类型的方法集满足接口定义所有方法，编译器自动认为该类型实现了该接口。
*/

// 定义动物接口
type Animal interface {
	// 定义动物行为
	Eat(food string) // 吃
	Speak()          // 叫
	// 定义动物属性
	Info()
}

// 定义狗及其实现方法
type Dog struct {
	Name string
}

func (d Dog) Eat(food string) {
	fmt.Printf("狗吃了%s\n", food)
}

func (d Dog) Speak() {
	fmt.Printf("汪汪汪！\n")
}

func (d Dog) Info() {
	fmt.Printf("名字：%s\n", d.Name)
}

// 继续定义猫接口
type Cat interface {
	// 组合形式“继承”动物接口
	Animal
	// 猫独立的行为
	SharpenClaws() // 磨爪
}

// 狸花猫想实现Cat接口时需要实现全部方法，包括内部的Animal接口
type TabbyCat struct {
	Name string
}

func (t TabbyCat) Eat(food string) {
	fmt.Printf("狸花猫吃了%s\n", food)
}

func (t TabbyCat) Speak() {
	fmt.Printf("喵喵喵！\n")
}

func (t TabbyCat) Info() {
	fmt.Printf("名字：%s\n", t.Name)
}

func (t TabbyCat) SharpenClaws() {
	fmt.Printf("狸花猫磨爪\n")
}

// 空接口测试
func TestEmptyInterface(t *testing.T) {
	// 空接口可以接受任何值，因为空接口什么方法都没有，因此任意类型都隐式实现了空接口
	var emptyInterface any // 或按旧版定义成 var emptyInterface interface{}
	// 常见值类型
	emptyInterface = 1                 // 整数
	emptyInterface = 1.0               // 浮点
	emptyInterface = "hello Interface" // 字符串
	emptyInterface = Dog{Name: "旺财"}   // 结构体
	// 引用类型
	emptyInterface = make([]int, 5)              // 切片
	emptyInterface = make(map[string]string, 10) // Map
	emptyInterface = make(chan int)              // 通道
	emptyInterface = &Dog{Name: "旺财"}            // 指针
	// 使用前需要进行断言。断言时尽量使用双值返回中ok，避免类型不匹配，程序出现panic
	if d, ok := emptyInterface.(*Dog); ok {
		d.Speak()
	}
}

// 多重接口测试
func TestInterface(t *testing.T) {
	var animal Animal
	var cat Cat
	var dog = Dog{"旺财"}
	animal = cat // 传入子接口时，不报错
	// animal.Info() // 但使用方法时，子接口只定义方法，但是没有提供方法，即animal接口中类型有，但是值为nil
	animal = dog  // 传入实现接口的类型，不报错
	animal.Info() // 使用方法，由于值和类型都满足，故可以调用

	var lihua = TabbyCat{"狸花"}
	cat = lihua // 同上，传入实现接口的具体类型
	cat.SharpenClaws()
	animal = cat // animal拿到的时cat此时持有的底层的类型和值，能正常调用所有方法，但不包括cat接口的方法SharpenClaws()
	animal.Info()
	animal = lihua
	animal.Speak()
	// 简单来说只要方法签名全满足，即为实现对应接口
}

package factory

import "fmt"

type FactoryAbstract interface {
	Greet()
}

// 不可导出的结构体，不公开内部实现
type factoryabstract struct {
	name string
	age  int
}

// 实现FactoryAbstract
func (f factoryabstract) Greet() {
	fmt.Printf("Hi! My name is %s", f.name)
}

// 返回一个 interface 隐藏内部实现
func newFactroyAbastract(name string, age int) FactoryAbstract {
	return factoryabstract{
		name: name,
		age:  age,
	}
}

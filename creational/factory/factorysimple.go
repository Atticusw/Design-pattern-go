package factory

type factorysimple struct {
	Name string
	Age  int
}

// 简单工厂模式
func (f factorysimple) newFactorySimple(name string, age int) *factorysimple {
	return &factorysimple{
		Name: name,
		Age:  age,
	}
}

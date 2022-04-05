package singleton

type singletonhungry struct {
}

// 包被加载时创建
var ins *singletonhungry = &singletonhungry{}

func GetInsOr() *singletonhungry {
	return ins
}

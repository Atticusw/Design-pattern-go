package singleton

type singletonlazy struct {
}

var insLazy *singletonlazy

// 非并发安全
func GetInsOrLazy() *singletonlazy {
	if ins == nil {
		insLazy = &singletonlazy{}
	}
	return insLazy
}

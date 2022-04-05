package singleton

import "sync"

type singletononce struct {
}

var insonce *singletononce

// once.Do 保证 insonce 全局只被创建一次
var once sync.Once

func GetInsOnce() *singletononce {
	once.Do(func() {
		insonce = &singletononce{}
	})

	return insonce
}

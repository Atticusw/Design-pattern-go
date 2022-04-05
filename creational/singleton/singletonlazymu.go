package singleton

import "sync"

type singletonlazymu struct {
}

var insMu *singletonlazymu
var mu sync.Mutex

// 加锁，保证线程安全
func GetInsMu() *singletonlazymu {
	if insMu == nil {
		mu.Lock()
		if insMu == nil {
			insMu = &singletonlazymu{}
		}
		mu.Unlock()
	}
	return insMu
}

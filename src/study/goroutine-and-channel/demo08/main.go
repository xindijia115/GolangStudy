package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var rwm sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(i int) {
			rwm.RLock()
			// 各个读锁之间不必等待另一个锁释放
			fmt.Println("Ready Lock reading i:", i)
			time.Sleep(time.Second * 2)
			rwm.RUnlock()
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	// 读锁没有完全释放 写 这个进程就会一直等待
	// 需要等全部的读锁释放
	rwm.Lock()
	fmt.Println("Ready Locked writing ")
	rwm.Unlock()
}

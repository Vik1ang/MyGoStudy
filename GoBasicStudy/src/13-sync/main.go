package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	l1 := &sync.Mutex{}
	go lockFun1(l1)
	go lockFun1(l1)
	go lockFun1(l1)
	go lockFun1(l1)

	fmt.Println("================================")

	l2 := &sync.RWMutex{}
	go lockFun2(l2)
	go lockFun2(l2)
	go lockFun2(l2)
	go lockFun2(l2)

	fmt.Println("================================")

	for {
	}

}

func lockFun1(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("疯狂刮痧")
	time.Sleep(time.Second)
	lock.Unlock()
}

func lockFun2(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("疯狂治疗")
	time.Sleep(time.Second)
	lock.RUnlock()
}

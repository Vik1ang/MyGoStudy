package main

import (
	"fmt"
	"sync"
)

func SyncClass() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println(i)
		})
	}
}

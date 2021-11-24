package main

import (
	"fmt"
	"sync"
)

func main() {
	//fmt.Println("================================")
	//go Run1()
	//time.Sleep(1 * time.Second)

	//fmt.Println("================================")

	/* go Run1()
	i := 0
	for i < 100 {
		i++
		fmt.Println(i)
	} */

	/* var wg sync.WaitGroup
	wg.Add(1)
	go Run2(&wg)
	wg.Wait() */

	fmt.Println("================================")
	// channel
	c1 := make(chan int, 1)
	c1 <- 1
	fmt.Println(<-c1)

	fmt.Println("================================")

	c2 := make(chan int)

	go func() {
		c2 <- 1
	}()

	fmt.Println(<-c2)

	fmt.Println("================================")

	c3 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c3 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c3)
	}

	fmt.Println("================================")

	c4 := make(chan int, 5)
	var readc <-chan int = c4
	var writec chan<- int = c4
	writec <- 1
	fmt.Println(<-readc)

	fmt.Println("================================")
	c5 := make(chan int, 5)
	c5 <- 1
	c5 <- 2
	c5 <- 3
	c5 <- 4
	c5 <- 5
	close(c5)
	fmt.Println(<-c5)
	fmt.Println(<-c5)
	fmt.Println(<-c5)
	fmt.Println(<-c5)
	fmt.Println(<-c5)

	fmt.Println("================================")

	c6 := make(chan int, 1)
	c7 := make(chan int, 1)
	c8 := make(chan int, 1)

	c6 <- 1
	c7 <- 1
	c8 <- 1

	select {
	case <-c6:
		fmt.Println("c6")
	case <-c7:
		fmt.Println("c7")
	case <-c8:
		fmt.Println("c8")
	default:
		fmt.Println("都不满足")
	}

	fmt.Println("================================")

	c9 := make(chan int)
	var readc9 <-chan int = c9
	var writec9 chan<- int = c9
	go SetChan1(writec9)

	GetChan1(readc9)

}

func Run1() {
	fmt.Println("我跑起来了")
}

func Run2(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}

func SetChan1(writec chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("我是在Set函数")
		writec <- i
	}
}

func GetChan1(readc <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("我是在Get函数, 取出SetChan1返回%v\n", <-readc)

	}
}

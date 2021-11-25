package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	flag := make(chan bool)
	message := make(chan int)

	go son(flag, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	flag <- true
	time.Sleep(time.Second)
	fmt.Println("主线程结束")

	fmt.Println("================================")

	ctx, clear := context.WithCancel(context.Background())
	message1 := make(chan int)
	go son1(ctx, message1)
	for i := 0; i < 10; i++ {
		message1 <- i
	}
	clear()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

	fmt.Println("================================")

	ctx2_ := context.WithValue(context.Background(), "name", "qm")
	ctx2, clear2 := context.WithCancel(ctx2_)
	message2 := make(chan int)
	go son2(ctx2, message2)
	for i := 0; i < 10; i++ {
		message2 <- i
	}
	clear2()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

	fmt.Println("================================")

	ctx3, clear3 := context.WithDeadline(context.Background(), time.Now().Add(time.Second*8))
	message3 := make(chan int)
	go son3(ctx3, message2)
	for i := 0; i < 10; i++ {
		message3 <- i
	}
	defer clear3()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

	fmt.Println("================================")

	ctx4, clear4 := context.WithTimeout(context.Background(), time.Second*8)
	message4 := make(chan int)
	go son4(ctx4, message2)
	for i := 0; i < 10; i++ {
		message4 <- i
	}
	defer clear4()
	time.Sleep(time.Second)
	fmt.Println("主进程结束")

}

func son(flag chan bool, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值, %d\n", m)
		case <-flag:
			fmt.Println("我结束了")
			return
		}
	}
}

func son1(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值, %d\n", m)
		case <-ctx.Done():
			fmt.Println("我结束了")
			return
		}
	}
}

func son2(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值, %d\n", m)
		case <-ctx.Done():
			fmt.Println("我结束了", ctx.Value("name"))
			return
		}
	}
}

func son3(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值, %d\n", m)
		case <-ctx.Done():
			fmt.Println("我结束了")
			return
		}
	}
}

func son4(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到值, %d\n", m)
		case <-ctx.Done():
			fmt.Println("我结束了")
			return
		}
	}
}

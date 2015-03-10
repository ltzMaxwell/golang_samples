package main

import (
	"fmt"
	"math/rand"
	"time"
)

func do_stuff(x int) int { // 一个比较耗时的事情，比如计算
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模拟计算
	return 100 - x                                              // 假如100-x是一个很费时的计算
}

func branch(x int) chan int { // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
	ch := make(chan int)
	go func() {
		ch <- do_stuff(x)
		fmt.Println("got data from stuff")
	}()
	fmt.Println("returning ch from branch")
	return ch
}

func fanIn(branches ...chan int) chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < len(branches); i++ { //select会尝试着依次取出各个信道的数据
			select {
			case v1 := <-branches[i]:
				ch <- v1
			}
		}
	}()

	return ch
}

func main() {
	result := fanIn(branch(1), branch(2), branch(3))
	fmt.Println("got results")
	for i := 0; i < 3; i++ {
		// fmt.Println("going to print")
		fmt.Println("result is :", result)
		msg := <-result
		fmt.Println(msg)
	}
}

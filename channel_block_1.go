package main

import "fmt"

import "time"

func main() {

	channel := make(chan string) //注意: buffer为1

	go func() {
		fmt.Println("goroutine")
		// time.Sleep(1 * time.Second)
		fmt.Println("writing hello")
		channel <- "hello"
		fmt.Println("write \"hello\" done!")

		time.Sleep(1 * time.Second)
		fmt.Println("writing world")
		channel <- "World" //Reader在Sleep，这里在阻塞
		fmt.Println("write \"World\" done!")
	}()

	fmt.Println("before sleep in main")
	time.Sleep(1 * time.Second)
	msg := <-channel
	fmt.Println("Reader: ", msg)
	//block read
	fmt.Println("wating for data to read")
	msg = <-channel
	fmt.Println("Reader: ", msg)

}

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("hello")
	fmt.Printf("length of data is %d \n", len(data))
	// fmt.Println("data is % s", data)
	result := sha1.Sum(data)
	fmt.Printf("% x \n", result)
	fmt.Println("string is :", string(result[:]))
}

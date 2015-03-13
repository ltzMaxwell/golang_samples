package main

import (
	"fmt"
)

func main() {
	data := []byte{'A', 'C', 'C', 'B', 'A'}
	fmt.Println("slice length is :", len(data))
	data = append(data, 'L', 'M')
	fmt.Println("length after append is :", len(data))

	test0 := []int{1, 2, 3, 4, 5}
	test1 := make([]int, 5, 10)
	fmt.Println("length of test1 is :", len(test1))
	fmt.Println("capacity of test1 is :", cap(test1))
	copy(test1, test0)
	fmt.Println("length of test1 after copy is :", len(test1))
	for _, v := range test1 {
		fmt.Println("item in test1 is :", v)
	}
}

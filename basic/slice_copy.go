package main

import "fmt"

func main() {
	s := [5]int{1, 2, 3, 4, 5}

	s1 := s[1:3]
	s2 := s[0:2]
	fmt.Println(s1, s2)

	//  s1和s2指向同一个底层数组，copy只是数据上的变化，而没有影响到各个切片的指向位置！
	fmt.Println("s2 before copy is :", s2)
	at := copy(s2, s1)
	fmt.Println("s2 is :", s2)
	fmt.Println("at is :", at)
	fmt.Println(s, s1, s2)

	s1[0] = 9
	fmt.Println(s, s1, s2)
}

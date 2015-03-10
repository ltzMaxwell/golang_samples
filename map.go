package main

import (
	"fmt"
)

func main() {
	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	fmt.Println("value of C is :", rating["C"])
	delete(rating, "C") // 删除key为C的元素
	fmt.Println("value of C after delete is :", rating["C"])
}

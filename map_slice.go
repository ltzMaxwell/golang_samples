package main

import "fmt"

func sumPointerToArray(a *[8]int) (sum int) {
	for _, value := range *a {
		sum += value
	}
	return
}
func sumSlice(a []int) (sum int) {
	// fmt.Println("element passed from slice is :", a)
	fmt.Println("a 0 is :", a[0])
	for _, value := range a {
		sum += value
	}
	return
}
func main() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := []int{1, 2, 3, 4}
	fmt.Printf("sum arrray via pointer: %d\n", sumPointerToArray(&array))
	fmt.Printf("sum slice: %d\n", sumSlice(slice))
	slice = array[0:]
	fmt.Printf("sum array as slice: %d\n", sumSlice(slice))
}

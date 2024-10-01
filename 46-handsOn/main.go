package main

import "fmt"

func main() {
	// to delete from a slice we use append and slicing
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}

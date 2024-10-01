package main

import "fmt"

func main() {
	si := []int{42, 43, 44, 45, 46, 47, 48, 49}
	for _, v := range si {
		fmt.Printf("%v-%T \n", v, v)
	}

}

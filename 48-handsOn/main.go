package main

import "fmt"

func main() {
	x1 := []string{
		"James",
		"bond",
		"Shaken, not started",
	}
	x2 := []string{
		"Miss",
		"Moneypenny",
		"I'm 008.",
	}
	ss := [][]string{x1, x2}
	//fmt.Println(ss)
	for i, v := range ss {
		fmt.Println(i, v)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}
}

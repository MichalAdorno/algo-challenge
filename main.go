package main

import "fmt"

func main() {
	var v []int
	fmt.Println(v == nil && len(v) == 0)
}

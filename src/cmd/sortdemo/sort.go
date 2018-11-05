package main

import (
	"fmt"
	"log"
)

func main() {
	//a := []int {3,6,2,1,9,10,8}
	//sort.Ints(a)
	//
	//for _,v := range a {
	//	fmt.Println(v)
	//}
	i := a(25)
	fmt.Print(i)
}

func a(i int64) int64 {
	if i < 0 {
		log.Panic("i must not be negative int")
	}
	if i == 1 || i == 0 {
		return 1
	}
	return i * a(i-1)
}

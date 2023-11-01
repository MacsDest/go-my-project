package main

import (
	"fmt"
	"go-my-project/cmd/my-app01"
	"math"
)

func main() {
	my_app01.Sum()
	fmt.Println(math.MinInt8)
	arr := [10]int{}
	fmt.Println(arr)
	arr1 := arr
	arr1[7] = 99
	fmt.Println(arr1)
	fmt.Println(arr)
	a := [...]int{99: 1}
	var p *[100]int = &a
	fmt.Println(*p)

}

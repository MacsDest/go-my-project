package main

import "fmt"

func main() {
	arr := [...]int{5, 4, 23, 54, 7, 2, 9}
	fmt.Println(arr)
	num := len(arr)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr[i] > arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	fmt.Println(arr)
}

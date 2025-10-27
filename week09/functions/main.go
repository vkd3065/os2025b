package main

import "fmt"

func swap(first int, second int) {
	temp := first
	first = second
	second = temp
	fmt.Println(first, second)

}

func mian() {
	a, b := 10, 20
	fmt.Println(a, b)
	swap(a, b)
	fmt.Println(a, b)
}

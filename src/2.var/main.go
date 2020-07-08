package main

import "fmt"

var globalA = 5

func main() {
	var a string = "goorm"
	fmt.Println(a)

	var b int = 10
	fmt.Println(b)

	var d = true //변수를 지정할 때 변수 명이 먼저 나오고 다음으로 타입이 나온다.
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	fmt.Println(globalA)
}

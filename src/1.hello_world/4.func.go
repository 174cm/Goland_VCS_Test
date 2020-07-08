package main

// 매개변수(인자)를 가지는 함수(1)
import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

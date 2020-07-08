package main

//하나의 함수는 여러 개의 결과를 반환할 수 있음.
import "fmt"

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("hello", "bye")
	fmt.Println(a, b)
}

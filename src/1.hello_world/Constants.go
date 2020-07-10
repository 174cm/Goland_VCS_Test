package main

// 상수는 const 키워드와 함께 변수처럼 선언함. 이상한건 const는 뒤에 안붙이네.
// 상수는 char, string, boolean 숫자 타입 중의 하나가 될 수 있음.

import "fmt"

const Pi = 3.14

func main() {
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

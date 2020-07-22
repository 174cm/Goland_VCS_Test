/*
- 요구 사항
1. var 키워드와 int 자료형을 모두 명시한 num1에 3을 선언 및 초기화
2. := 연산자를 사용해 num2에 7을 선언 및 초기화
3. := 연산자를 사용해 num1과 num2의 덧셈의 연산 및 결과 값을 result 변수에 선언 및 초기화
4. printf를 사용해 결과값을 출력

*/

package main

import "fmt"

func main() {
	var num1 int = 3
	num2 := 7
	result := num1 + num2

	fmt.Printf("%d입니다.", result)
}

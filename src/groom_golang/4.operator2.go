/*
- 요구 사항
1. 2개의 정수형 변수를 입력받고 몫과 나머지를 출력한다.
2. 예를 들어, 7과 2를 순서대로 입력했다면, 몫은 3, 나머지는 1입니다.
*/

package main

import "fmt"

var num1, num2 int

func main() {
	fmt.Printf("정수 2개를 입력하세요: ")
	fmt.Scanln(&num1, &num2)
	fmt.Printf("몫 : %d, 나머지 : %d", num1/num2, num1%num2)
}

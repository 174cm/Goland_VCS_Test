/*
- 요구 사항
1. 3개의 정수와 결과값을 저장할 변수 num1, num2, num3, result를 선언한다.
2. Scanln()함수로 사용자로부터 num1, num2, num3를 입력 받는다.
3. num1xnum2+num3 결과값을 result에 초기화한다.
4. num1, num2, num3, result 값을 Printf 함수를 사용해 출력
*/

package main

import "fmt"

func main() {
	var num1, num2, num3, result int

	fmt.Println("정수 3개를 입력하세요: (띄어쓰기로 인식함)")
	fmt.Scanln(&num1, &num2, &num3)
	result = num1*num2 + num3
	fmt.Printf("%d x %d + %d = %d\n", num1, num2, num3, result)

}

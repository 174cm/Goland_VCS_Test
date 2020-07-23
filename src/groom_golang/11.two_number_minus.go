/*
- 요구 사항
1. int형 변수 num1, num2, result 선언
2. num1과 num2 연산에 사용되고 result는 결과값을 저장합니다.
3. 두 정수를 입력 받아서 두 수의 차를 출력하는 프로그램 작성
4. 이때, 무조건 큰 수에서 작은 수를 뺀 결과를 출력해야 함.
*/
package main

import "fmt"

func main() {
	var num1, num2, result int
	fmt.Printf("두 정수를 입력하세요: ")
	fmt.Scan(&num1, &num2)

	if num1 < num2 {
		result = num2 - num1
		fmt.Println(result)

	} else if num1 > num2 {
		result = num1 - num2
		fmt.Println(result)
	}
}

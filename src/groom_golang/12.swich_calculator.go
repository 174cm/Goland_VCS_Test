/*
- 요구 사항
1. 사용자에게 입력받을 연산 번호를 저장할 정수형 변수 sel을 선언
2. 연산을 할 num1과 num2, 결과값을 저장할 result를 실수형으로 선언
3. 1번은 덧셈, 2번은 뺼셈, 3번은 곱셈, 4번은 나눗셈을 연산함.
4. 이외의 숫자를 입력할 시 "잘못된입력입니다."를 출력하고 프로그램이 종료됨.
*/
package main

import "fmt"

func main() {
	var sel int
	var num1, num2, result float32

	fmt.Scanln(&sel)         //연산 번호 입력
	fmt.Scanln(&num1, &num2) //연산할 두 변수 입력

	switch sel {
	case 1: //덧셈
		result = num1 + num2

	case 2: //뺼셈
		result = num1 - num2

	case 3: //곱셈
		result = num1 * num2

	case 4: //나눗셈
		result = num1 / num2

	default:
		fmt.Println("잘못된입력입니다.")
		return
	}
	fmt.Printf("%.1f\n", result)
}

/*
- 요구 사항
1. int형 변수 num1, num2, num3를 선언
2. 사용자로부터 입력받을 때 두 번째 수는 무조건 '음수'
3. num1은 float32형으로 num2는 uint형, num3는 int64형으로 변환 후, 새로운 변수에 초기화
4. Printf 함수를 사용해 형 변환 된 변수 세 개가 출력됩니다.
5. 타입을 출력하는 서식문자는 '%T', 실수형을 출력하는 서식문자는 '%f'를 입력

*/
package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("정수 3개를 입력하세요: ")
	fmt.Scanln(&num1, &num2, &num3)
	result1 := float32(num1)
	result2 := uint(num2)
	result3 := int64(num3)

	fmt.Printf("%T, %f\n", result1, result1)
	fmt.Printf("%T, %d\n", result2, result2)
	fmt.Printf("%T, %d\n", result3, result3)

}

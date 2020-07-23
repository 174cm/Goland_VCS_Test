/*
- 요구 사항
1. 단 수를 입력받을 int형 변수 dan을 선언
2. 사용자로 부터 dan을 입력받음
3. 7 x 3 = 21형태로 출력
4. 1부터 9까지 곱셈을 출력
*/

package main

import "fmt"

func main() {
	var dan, n int
	n = 1
	fmt.Println("구구단을 입력해주세요:")
	fmt.Scanln(&dan)
	for n < 10 {
		fmt.Printf("%d x %d = %d\n", dan, n, dan*n)

		n++
	}

}

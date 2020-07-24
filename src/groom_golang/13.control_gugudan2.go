/*
- 요구 사항
1. 2단부터 9단까지 출력하는 구구단. 단, 홀수 다만 출력함.
2. "%d x %d = %d\n"형태로 출력
3. 단과 단 사이는 한 줄을 비움
4. n단은 n x n 까지 출력. 예를 들어, 7단은 7 x 7 = 49까지 출력
*/

package main

import "fmt"

func main() {

	for i := 2; i < 10; i++ {
		if i%2 == 0 { // 홀수 출력
			continue

		}
		for dan := 1; dan < 10; dan++ { // 단 수 출력
			if dan > i { // n단은 n x n까지 출력
				continue
			}

			fmt.Printf("%d x %d = %d\n", i, dan, i*dan)
		}
		fmt.Printf("\n") //단 마다 한칸 비움
	}
}

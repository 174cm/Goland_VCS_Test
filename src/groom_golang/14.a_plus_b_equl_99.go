/*
- 요구 사항
1. 10의 자리가 A이고 1의 자리가 B인 두 자릿수와, 10의 자리가 B이고 1의자리가 A인 두 자릿수가 더해서 99가 나오는 모든 수를 출력
2. A와 B는 다른 숫자, 따라서 33 + 66 = 99는 잘못된 출력
3. 가능한 모든 경우의 수를 출력함
4. 이중 반복문을 사용
5. 1의자리에 0을 출력할 수 있음.
*/
package main

import "fmt"

func main() {
	var result int

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			result = (a*10 + b) + (a + b*10) // 10의자리 출력

			if result != 99 {
				continue
			}
			fmt.Printf("%d%d+%d%d = %d\n", a, b, a, b, result) //출력형식지정.
		}
	}
}

/**
- 요구 사항
1. 1이상 100미만의 정수 중에서 7의 배수와 9의 배수를 출력하는 프로그램을 작성
2. 단, 7의 배수이면서 동시에 9의 배수인 정수는 한 번만 출력한다.
*/

package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)

		} else if i%9 == 0 {
			fmt.Printf("%d ", i)
		}
		//fmt.Printf("%d ", i)
	}
}

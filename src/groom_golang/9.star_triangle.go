/*
- 요구 사항
1. 사용자로부터 이등변삼각형의 빗변을 제외한 같은 값의 두 변의 길이를 입력받음.
2. 빗면이 * 모양으로 빛나는 이등변 삼각형이 출력
3. 기호와 기호 사이는 띄어쓰기를 사용.
*/
package main

import "fmt"

func main() {
	var i, j, k int
	//i,j는 두 개의 반복문에 쓰일 변수
	fmt.Printf("별모양 이등변 삼각형을 만들 수를 입력해주세요: ")
	fmt.Scanln(&k)

	for i = 0; i < k; i++ {
		for j = 0; j < i; j++ {
			fmt.Print("o ")
		}
		fmt.Println("* ")
	}
}

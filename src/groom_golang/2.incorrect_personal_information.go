/*
- 요구 사항
1. main 함수 밖에 const 변수 name, RRN, job을 순서대로 괄호를 사용해 묶어서 선언
2. name은 kim, RRN은 800101-1000000으로 초기화. 모두 문자열
3. job은 초기화하지 않고 선언만 명시
4. main 함수 내에 Println 함수를 사용해 name, RRN, job을 순서대로 출력
(RRN:주민번호)
*/

package main

import "fmt"

const ( //괄호로 묶어서 정의해야 함.
	name = "kim"
	RRN  = "800101-1000000"
	job
)

func main() {

	fmt.Println(name, RRN, job)
}

/*
- 요구 사항
1. 주민등록번호 앞자리와 뒷자리를 저장할 int형 변수 RRNf와 RRNt를 선언
2. 이름을 저장할 string형 변수 name을 선언
3. 키를 저장할 float32형 변수 height를 선언
4. 첫 번째 줄부터 주민등록번호, 이름, 키를 입력받음
5. 정보를 모두 입력하면 결과값이 출력됨
*/

package main

import "fmt"

func main() {
	var RRNf, RRNt int
	var name string
	var height float32

	fmt.Scanf("%d-%d", &RRNf, &RRNt)
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &height)

	fmt.Printf("주민등록번호 앞자리는 %d, 뒷자리는 %d, 이름은 %s입니다.\n", RRNf, RRNt, name)
	fmt.Printf("그리고 키는 %.2f입니다.", height)
}

/*
- 요구 사항
1. 이름 열은 전부 폭을 8로 지정, 왼쪽 정렬
2. 전공학과 열은 전부 폭을 14로 지정, 왼쪽 정렬
3. 학년 열은 전부 폭을 5로 지정, 오른쪽 정렬
4. 모든 값은 string 형

*/

package main

import "fmt"

func main() {
	fmt.Printf("%-8s%-14s%5s\n", "이름", "전공학과", "학년")
	fmt.Printf("%-8s%-14s%5s\n", "유현수", "전자공학", "3")
	fmt.Printf("%-8s%-14s%5s\n", "김윤욱", "컴퓨터공학", "4")
	fmt.Printf("%-8s%-14s%5s\n", "김나영", "미술교육학", "2")
}
//역행렬
package main

import "fmt"

func main() {
	A := [2][2]float32{ //2차원 행렬 선언
		{7, 3}, {5, 2}, //1행에 7,3 2행에 5,2 원소
	}
	Invert_A := [2][2]float32{} //배열 초기화
	var B bool = true

	d := A[0][0]*A[1][1] - A[0][1]*A[1][0] //역행렬 식.

	if d != 0 { //0이 아니면 행렬 A를 역행렬로 만드는 연산을 수행. 그리고 TRUE와 연산 결과인 역행렬을 출력함.
		Invert_A[0][0] = A[1][1] / d
		Invert_A[0][1] = A[0][1] / -d
		Invert_A[1][0] = A[1][0] / -d
		Invert_A[1][1] = A[0][0] / d
		B = true
		fmt.Print(B)
		fmt.Printf("\n")

		for i := 0; i < 2; i++ { //연산 수행
			for j := 0; j < 2; j++ {
				fmt.Print(Invert_A[i][j])
				fmt.Printf(" ")
			}
			fmt.Println("")
		}

	} else { // 0이면 아무 연산도 하지 않고 False만 출력함.
		B = false
		fmt.Print(B)
	}
}

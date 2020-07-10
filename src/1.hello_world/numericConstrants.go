package main

//숫자형 상수는 정밀한 값을 표현할 수 있다.
//타입을 지정하지 않은 상수는 문맥에 따라 타입을 가지게 된다.

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

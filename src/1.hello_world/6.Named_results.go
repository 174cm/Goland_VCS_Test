package main

// 반환 값에 이름을 취함으로써 변수처럼 사용이 가능함.
import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(split(20))
}

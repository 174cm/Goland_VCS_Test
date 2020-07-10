package main

// go에서는 반복문이 for 밖에 존재하지 않으며 c, java와 모습이 거의 동일
//소괄호가 필요하지 않다는 점과 대괄호로 {} 실행문을 지정함.
import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

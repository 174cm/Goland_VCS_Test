package main

// :=을 사용하면 var과 명시적인 타입 (int bool)을 생략 가능
// 하지만 함수 밖에서는 := 선언을 할 수 없음.
import "fmt"

func main() {
	var x2, y2, z2 int = 1, 2, 3
	c2, python2, java2 := true, false, "no1"

	fmt.Println(x2, y2, z2, c2, python2, java2)
}

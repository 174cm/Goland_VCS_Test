package main

//조건문만 표시하여 for문을 돌릴 수 있음.
//해당 표현식은 while을 사용하듯이 for를 사용할 수 있음
import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

package main

// 한눈에 끝내는 고랭 기초의 struct
import "fmt"

type person struct {
	name, contact string
	age           int
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "sungho"
	p1.age = 27
	p1.contact = "01011112222"
	fmt.Println(p1)

	p2 := person{"min", "01055553333", 33} //필드 이름을 생략 시, 순서대로 저장함
	fmt.Println(p2)

	p3 := person{contact: "01033334444", name: "sanghwa", age: 90} // 필드 이름을 명시할 시, 순서와 상관 없이 저장할 수 있음
	fmt.Println(p3)

	p3.name = "ryu" // 필드에 저장된 값을 수정할 수 있음
	fmt.Println(p3)

	fmt.Println(p3.contact) // 필드 값의 개별 접근도 가능함
}

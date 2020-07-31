package main

import "fmt"

//map은 hash table을 지원하는데 key:value 값으로 저장한다. 사전과 같은 방식.
func main() {
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "충청도"
	m["053"] = "대구"

	fmt.Println(m)

	m["033"] = "인천"

	fmt.Println(m)

	delete(m, "031")

	fmt.Println(m)
}

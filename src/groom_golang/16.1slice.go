package main

import "fmt"

/*
//1. 슬라이스의 길이와 용량을 지정하지 않고 슬라이스를 선언만 해서 Nil slice를 만들어 nil과 비교할 수 있고 true를 반환함.
func main() {
	var a []int        // 슬라이스 변수 선언. 초기화 되지 않은 상태
	a = []int{1, 2, 3} // 슬라이스에 리터럴값 지정

	a[1] = 10 //값이 할당되어 메모리가 생겼기에 다음과 같이 접근 가능.

	fmt.Println(a)

	var b []int //nil slice 선언

	if b == nil {
		fmt.Println("용량이", cap(b), "길이가", len(b), "NIL SLICE입니다.")
	}
}
*/

/*
// 2. make()를 이용한 선언. make(slice type, slice len, slie cap)
// slice len: 초기화된 슬라이스의 요소 갯수. 즉, 슬라이스에 5개의 값이 초기화된다면 길이는 5가 됨. 그 후에 값을 추가하거나 삭제한다면 그만큼 길이가 바뀜.
// slice cap(capacity): 슬라이스는 배열의 길이가 동적으로 늘어날 수 있기 때문에 길이와 용량을 구분함.
// cap의 내용을 이해하기가 어려웠는데, 용량이 다 차면 자동으로 두배 씩 늘려감. 아래의 예에서는 3, 6, 12로 늘어나는 것을 볼 수 있음.

func main()  {
	s := make([]int,0,3)

	for i := 1; i <= 10; i++ {
		s = append(s, i)

		fmt.Println(len(s), cap(s))
	}

	fmt.Println(s)
}
*/

/*
//3. 슬라이스 추가, 병합, 복사

func main() {
	sliceA := []int{1, 2, 3} //a 슬라이스 선언
	sliceB := []int{4, 5, 6} //b 슬라이스 선언

	sliceA = append(sliceA, sliceB...) //a 슬라이스에 {4,5,6} 요소를 추가함.

	fmt.Println(sliceA)
}
*/

//copy
/*
func main() {
	sliceA := []int{0, 1, 2}
	sliceB := make([]int, len(sliceA), cap(sliceA)*2) //슬라이스 b는 a의 2배 용량으로 선언

	copy(sliceB, sliceA) //copy(붙여넣을 슬라이스, 복사할 슬라이스), a를 b에 붙여 넣음.

	fmt.Println(sliceB) //[0.1.2]
	println(len(sliceB), cap(sliceB)) //길이가 3, 용량이 6 출력
}
*/

func main() {

	c := make([]int, 0, 3)
	fmt.Println(len(c), cap(c))
	c = append(c, 1, 2, 3, 4, 5, 6) // for loop는 컴파일러가 변수를 미리 인식하여 미리 2배씩 늘려가는 느낌이었고
	// 다음의 append는 미리 선언된 값을 사용하다 보니 +1정도로 여유를 둔 듯 함. (정확한 이유는 모름)
	fmt.Println(len(c), cap(c))

	l := c[1:3] //1~2요소까지 복사
	fmt.Println(l)

	l = c[2:] //2요소부터 끝까지 복사
	fmt.Println(l)

	l[0] = 6 // 바로 위출력으로 3,4,5,6,7로 초기화되어서 3이 0번째 요소가 됨.

	fmt.Println(c) // 따라서 3이 6으로 변경되었음.
}

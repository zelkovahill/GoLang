package main

import "fmt"

// 구조체
type Person struct {
	Name string
	Age  int
}

// 함수

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")

	var a int = 10
	var b = 20 // 타입 추론
	c := 30    // 단축 변수 선언

	fmt.Println(a, b, c)

	var f float32 = 3.14
	var B bool = true
	var s string = "Hello, World!"

	fmt.Println(f, B, s)

	// 정수 int , inm8, int16, int32, int64
	// 부동 소수점 : float32, float64
	// 불리언 : bool
	// 문자열 : string
	// 복소수 : complex64, complex128

	// 상수 const
	const pi = 3.14
	const Greating = "Hello, World!"

	fmt.Println(pi, Greating) // 3.14 Hello, World!

	//  제어문
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5") // x is greater than 5
	} else { // 엔터를 치면 안된다 else 가 붙어야 한다.
		fmt.Println("x is less than 5")
	}

	// 반복문
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}

	// 무한 루프
	// j := 0
	// for {
	// 	fmt.Println(j)
	// 	j++
	// }

	// 스위치 문
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday") // 3
	default:
		fmt.Println("Unknown")
	}

	// 함수는 위에 선언했음
	result := add(10, 20)
	fmt.Println(result) // 30

	// 배열과 슬라이스
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr) // [1 2 0 0 0]

	// 슬라이스
	w := []int{1, 2, 3}
	w = append(w, 4, 5)
	fmt.Println(w) // [1 2 3 4 5]

	// 맵
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m) // map[one:1 two:2]

	// 구조체
	l := Person{Name: "Lee", Age: 10}
	fmt.Println(l) // {Lee 10}

	// 인터페이스
	var a1 Animal

	a1 = Dog{}
	fmt.Println(a1.Speak()) // Woof!

	a1 = Cat{}
	fmt.Println(a1.Speak()) // Meow!

}

// 인터페이스
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

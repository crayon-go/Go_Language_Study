package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("name", nil)
	fmt.Println("Park", 10, 10, "word")

	st := User{"name", 10}

	// 구조체 출력
	fmt.Printf("%v\n", st)  // %v: 구조체 내용물 가져와서 출력
	fmt.Printf("%#v\n", st) // %#v: 해당 구조체를 호출하는 구간과 내용 출력 (디버깅용)
	fmt.Printf("%+v\n", st) // %+v: 해당 구조체의 변수 이름까지 함께 출력
	fmt.Printf("%T\n", st)  // $T: 해당 구조체 호출 구간 출력

	// 자료형 출력
	var num int = 10
	var name string = "Park"
	var byte1 byte = 'A'
	var boolean bool = true
	var f float32 = 1.5
	var p *int = &num

	fmt.Printf("-----------------\n")
	fmt.Printf("%d\n", num)     // 정수
	fmt.Printf("%f\n", f)       // 실수
	fmt.Printf("%s\n", name)    // 문자열
	fmt.Printf("%b\n", byte1)   // 2진수
	fmt.Printf("%0o\n", byte1)  // 8진수
	fmt.Printf("%x\n", byte1)   // 16진수
	fmt.Printf("%c\n", byte1)   // 문자
	fmt.Printf("%t\n", boolean) // boolean
	fmt.Printf("%U\n", byte1)   // 유니코드
	fmt.Printf("%p\n", &p)      // 포인터가 가르키는 주소 값
	fmt.Printf("%p\n", p)       // 포인터 주소 값

	// 문자열
	var name2 string = "Park"
	var str1 string = "abc\n" + "def"
	var str2 string = `abc\n` + name2
	var str3 string = `qqq\n 
	zzzz`

	fmt.Printf("-----------------\n")
	fmt.Println(`hello` + name2)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

}

package main

import "fmt"

func modify(s *string) {
	*s = "world"
}

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {
	a := "hello"
	fmt.Printf("%T %v \n", a, a)

	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

	var s Student
	fmt.Printf("%+v", s)

	st := new(Student)
	fmt.Println("")
	fmt.Printf("%+v", st)

	st1 := Student{"Joe", 12, []int{1, 2, 3}, map[string]int{"Jack": 40}}
	fmt.Println(st1)

}

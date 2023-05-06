package main

import "fmt"

func modify(s *string) {
	*s = "world"
}

func main() {
	s := "hello"
	fmt.Printf("%T %v \n", s, s)

	/*
		ps := &s
		*ps = "world"
		fmt.Printf("%T %v \n", s, s)
	*/

	fmt.Println(s)
	modify(&s)
	fmt.Println(s)

}

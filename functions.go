package main

import "fmt"

// Simple function adding two integer values
func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func operation(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Variadic functions
func printDetails(student string, subjects ...string) {
	fmt.Println("hey", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

func f() (int, int) {
	return 42, 53
}

func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

func main() {

	sumOfNumbers := addNumbers(10, 5)
	fmt.Println(sumOfNumbers)

	sum, diff := operation(10, 10)
	fmt.Println("Sum ->", sum, " Diff ->", diff)

	printDetails("Joe", "Physics", "Biology")

	// Blank identifier
	v, _ := f()
	fmt.Println()
	fmt.Println(v)

	// anonymous function
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))

	// Anonymous functions
	y := func(l int, b int) int {
		return l * b
	}(20, 30)

	fmt.Printf("%T \n", x)
	fmt.Println(y)

	// A defer statement delays the execution of a function until
	// the surrounding functions returns
	printName("Joe")
	defer printRollNo(23)
	printAddress("street-32")

}

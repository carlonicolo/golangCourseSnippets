package main

import "fmt"

func main() {
	var fruits [5]string = [5]string{"apples", "banana", "orange", "mango", "papaya"}
	fmt.Println(fruits)

	cars := [...]string{"BMW", "Mercedes", "FIAT", "VOLVO"}
	fmt.Println(cars)

	fmt.Println(len(cars))

	fmt.Println(cars[2])
	cars[2] = "Renault"

	fmt.Println(cars)

	for index, element := range cars {
		fmt.Println(index, " => ", element)
	}

	// Multidimensional array
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr)
	fmt.Println(arr[0][1])

}

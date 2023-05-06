package main

import "fmt"

func main() {
	codes := map[string]int{"en": 1, "fr": 2, "it": 3}
	value, found := codes["en"]
	fmt.Println(found, value)

	value, found = codes["zb"]
	fmt.Println(found, value)

	arr := [5]int{}
	my_map := make(map[string]int)
	my_map["A"] = 65
	my_map["B"] = 66
	i := 0
	for _, value := range my_map {
		arr[i] = value
		i += 1
	}
	fmt.Println(arr)

}

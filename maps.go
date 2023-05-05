package main

import "fmt"

func main() {
	codes := map[string]int{"en": 1, "fr": 2, "it": 3}
	value, found := codes["en"]
	fmt.Println(found, value)

	value, found = codes["zb"]
	fmt.Println(found, value)
}

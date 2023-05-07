package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printDetail(s shape) {
	fmt.Println("Struct -> ", s)
	fmt.Println("Area -> ", s.area())
	fmt.Println("Perimeter -> ", s.perimeter())
	fmt.Println("")
}

func main() {
	r := rect{length: 3, breadth: 4}
	c := square{side: 5}
	printDetail(r)
	printDetail(c)
}

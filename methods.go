package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

type Student struct {
	name   string
	grades []int
}

func (c *Circle) calcArea() {
	c.area = 3.14 * c.radius * c.radius
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v", c)

	fmt.Println("\n")
	s := Student{name: "Joe", grades: []int{90, 75, 80}}
	s.displayName()
	fmt.Printf("%.2f%%", s.calculatePercentage())
}

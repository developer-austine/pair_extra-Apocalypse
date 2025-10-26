package main

import "fmt"

type Calculations interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}

func main() {
	rectangle := Rectangle{
		Length: 15,
		Width:  20,
	}

	fmt.Println("Area:", rectangle.Area())
  fmt.Println("Perimeter:", rectangle.Perimeter())
}

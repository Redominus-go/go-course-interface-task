package main

import "fmt"

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		height: 4,
		base:   2,
	}
	squ := square{
		sideLength: 2,
	}
	printArea("triangle", tri)
	printArea("square", squ)
}

func printArea(n string, s shape) {
	fmt.Printf("Area of %v: %v\n", n, s.getArea())
}

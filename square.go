package square

//package main

import (
	"math"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	// implement me
	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}

func (s Square) Area() uint {
	// implement me
	return uint(math.Pow(float64(s.a), 2))
}

func (s Square) Perimeter() uint {
	// implement me
	return uint(4 * s.a)
}

//func main() {
//	var resPoint Point
//	var areaSquare, perimeterSquare uint
//
//	var sInput Square = Square{Point{1, 1}, 1}
//	resPoint = sInput.End()
//	areaSquare = sInput.Area()
//	perimeterSquare = sInput.Perimeter()
//	fmt.Printf("Begin point, X: %d, Y: %d, lenght of side: %d\n", sInput.start.x, sInput.start.y, sInput.a)
//	fmt.Printf("End point, X: %d, Y: %d\n", resPoint.x, resPoint.y)
//	fmt.Printf("Area square: %d\n", areaSquare)
//	fmt.Printf("Perimeter square: %d\n", perimeterSquare)
//}

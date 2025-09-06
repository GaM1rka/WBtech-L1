package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64 // Private field, because it is written with lowercase
	y float64 // Private field, because it is written with lowercase
}

func NewPoint(xVal, yVal float64) *Point {
	return &Point{
		x: xVal,
		y: yVal,
	}
}

// Getting value of private field `x`
func (p *Point) X() float64 {
	return p.x
}

// Getting value of private field `y`
func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Distance(other Point) float64 {
	res := math.Pow(other.X()-p.X(), 2) + math.Pow(other.Y()-p.Y(), 2)
	return math.Sqrt(res)
}

func main() {
	first := NewPoint(2.0, -5.0)
	second := NewPoint(-4.0, 3.0)
	fmt.Println(first.Distance(*second))
}

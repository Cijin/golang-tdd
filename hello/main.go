package main

import "math"

type Rectange struct {
	W float64
	H float64
}

type Circle struct {
	R float64
}

type Triangle struct {
	H float64
	B float64
}

type Shape interface {
	Area() float64
}

func perimeter(r Rectange) float64 {
	return 2 * (r.H + r.W)
}

func (r Rectange) Area() float64 {
	return r.H * r.W
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (t Triangle) Area() float64 {
	return 0.5 * t.H * t.B
}

// Interfaces project main.go
package main

import (
			"fmt"
			"math"
		)
		
// Structure for shape circle
type Circle struct{
	x, y, r float64
}
// Structure for shape rectangle
type Rectangle struct{
	x1, y1, x2, y2 float64
}

// Interface with a method set for area & perimeter 
type Shape interface{
	area() float64
	perimeter() float64
}

// Entry Function
func main(){
	
	// Interface shape
	var shape Shape
	
	// Circle with centre at (0,0) and radius 5
	cir := Circle{0,0,5}
	
	// Rectangle
	rect := Rectangle{0,0,3,6}
	
	// Circle Implementing interface Shape
	shape =  cir
	fmt.Println("Area of Cirlce : ",shape.area(),"square units.")
	fmt.Println("Perimeter of Circle : ",shape.perimeter(),"units.")
	
	// Rectangle implementing interface Shape
	shape = rect
	fmt.Println("Area of Rectangle : ", shape.area(),"square units.")
	fmt.Println("Perimeter of Rectangle : ", shape.perimeter(),"units.")
}

// Function to determine distance between two points
func distance(x1,y1,x2,y2 float64) float64{
	a := x2 - x1
	b := y2-y1
	return math.Sqrt(a*a + b*b)
}

// Function to determine area of a rectangle
func (r Rectangle) area() float64{
	l:= distance(r.x1,r.y1,r.x1,r.y2)
	w:= distance(r.x1,r.y1,r.x2,r.y1)
	return l*w
}

// Function to determine area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Function to determine perimeter of a rectangle
func (r Rectangle) perimeter() float64{
	l:= distance(r.x1,r.y1,r.x1,r.y2)
	w:= distance(r.x1,r.y1,r.x2,r.y1)
	return 2*(l+w)	
}

// Function to determine perimeter of a circle
func (c Circle) perimeter() float64{
	return 2 * math.Pi * c.r
}
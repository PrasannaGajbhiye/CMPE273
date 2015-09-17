package main

import "testing"

//struct for test pairs
type testPair struct{
	shape Shape
	perimeter float64
	area float64
}

// test cases
var tests = []testPair {
	//Pi=3.141592653589793
		{Circle{0.0,0.0,1.0}, 6.283185307179586, 3.141592653589793},
		{Rectangle{0,0,2,3}, 10, 6},
		{Rectangle{0,0,3,6}, 18, 18},
}

// Function to test perimeter function
func TestPerimeter(t *testing.T) {
	for _, pair:=range tests{
		var shp Shape
		shp=pair.shape
		v := shp.perimeter()
		if v != pair.perimeter {
	           t.Error(
					"For ", pair.shape,
					"expected ",pair.perimeter,
					"got ",v,
			)
    	}
	}
}

// Function to test area function
func TestArea(t *testing.T) {
	for _, pair:=range tests{
		var shp Shape
		shp=pair.shape
		v := shp.area()
		if v != pair.area {
	           t.Error(
					"For ", pair.shape,
					"expected ",pair.area,
					"got ",v,
			)
    	}
	}
}
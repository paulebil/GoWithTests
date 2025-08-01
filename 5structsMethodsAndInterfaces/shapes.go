package structsmethodsandinterfaces

import "math"

type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}


type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}


type Shape interface{
	Area() float64
}

type Triangle struct{
	Base float64
	Height float64
}

func (t Triangle) Area() float64{
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64{
	perimeter := 2 * (rectangle.Width + rectangle.Height)
	return perimeter
}

func Area(rectangle Rectangle) float64{
	area := rectangle.Width * rectangle.Height
	return area
}
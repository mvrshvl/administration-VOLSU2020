package geometry

import (
	"fmt"
	"math"
)

type triangle struct {
	x1,
	y1,
	x2,
	y2,
	x3,
	y3 float64
}

func NewTriangle(x1, y1, x2, y2, x3, y3 float64) triangle {
	return triangle{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
		x3: x3,
		y3: y3,
	}
}

func (t triangle) getArea() {
	s := 0.5 * math.Abs((t.x1-t.x3)*(t.y2-t.y3)-(t.x2-t.x3)*(t.y1-t.y3))
	fmt.Println("Triangle Area:", s)
}

func (t triangle) getPerimeter() {
	p := Length(t.x1, t.y1, t.x2, t.y2) + Length(t.x1, t.y1, t.x3, t.y3) + Length(t.x2, t.y2, t.x3, t.y3)
	fmt.Println("Triangle Perimeter:", p)
}

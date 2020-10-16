package geometry

import (
	"fmt"
	"math"
)

type triangle2 struct {
	x1,
	y1,
	x2,
	y2,
	x3,
	y3 float64
}

func NewTriangle2(x1, y1, x2, y2, x3, y3 float64) triangle2 {
	return triangle2{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
		x3: x3,
		y3: y3,
	}
}

func (t triangle2) Area() {
	s := 0.5 * math.Abs((t.x1-t.x3)*(t.y2-t.y3)-(t.x2-t.x3)*(t.y1-t.y3))
	fmt.Println("Triangle Area:", s)
}

func (t triangle2) Perimeter() {
	p := Length2(t.x1, t.y1, t.x2, t.y2) + Length2(t.x1, t.y1, t.x3, t.y3) + Length2(t.x2, t.y2, t.x3, t.y3)
	fmt.Println("Triangle Perimeter:", p)
}

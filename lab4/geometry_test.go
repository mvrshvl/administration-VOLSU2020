package lab4

import (
	"ADMINS/lab4/geometry"
	g2 "ADMINS/lab4/geometry2"
	"testing"
)

func TestGeometry(t *testing.T) {
	//неявно
	triangle := geometry.NewTriangle(0, 0, 1, 1, 2, 0)
	square := geometry.NewSquare(0, 0, 0, 1, 1, 1, 1, 0)
	geometry.Area(triangle)
	geometry.Area(square)
	geometry.Perimeter(triangle)
	geometry.Perimeter(square)
}

func TestGeometry2(t *testing.T) {
	var geometry g2.Geometry2

	triangle := g2.NewTriangle2(0, 0, 1, 1, 2, 0)
	geometry = &triangle
	geometry.Perimeter()
	geometry.Area()

	square := g2.NewSquare(0, 0, 0, 1, 1, 1, 1, 0)
	geometry = &square
	geometry.Perimeter()
	geometry.Area()

}

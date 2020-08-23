package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func TestGeometry() {
	r := rectangle{10, 20}
	c := circle{20}
	measure(r)
	measure(c)

}

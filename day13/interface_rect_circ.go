package interface4

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	parameter() float64
}

type rectangle struct {
	length  float64
	breadth float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) parameter() float64 {
	return 2 * (r.length + r.breadth)
}
func (c circle) parameter() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.parameter())
}

func main() {
	rect := rectangle{
		length:  10,
		breadth: 20,
	}
	cir := circle{
		radius: 5,
	}
	measure(rect)
	measure(cir)

}

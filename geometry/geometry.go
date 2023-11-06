package geometry

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func main() {

	p := Point{1, 2}
	q := Point{3, 4}

	fmt.Println(p.Distance(q))

	paths := Path{
		Point{2, 5},
		Point{3, 6},
		Point{4, 7},
	}
	fmt.Println("Paths Dist = %v", paths.Distance())
}

type Path []Point

func (paths Path) Distance() float64 {
	sum := 0.0
	for path_index := range paths {
		if path_index > 0 {
			sum += paths[path_index].Distance(paths[path_index-1])
			fmt.Println("Temp sum = %v", sum)
		}
	}
	return sum
}

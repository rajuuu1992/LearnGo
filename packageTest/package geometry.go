package geometry


import (
	"math"
)

type Point struct {
	X, Y float64
}

type
func (Point p) Distance (Point q) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func main() {

	p Point{1, 2}
	q Point{3,4}

	fmt.Println(p.Distance(q))
}
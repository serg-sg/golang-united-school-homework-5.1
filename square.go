package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func NewSquare(xx, yy int, a uint) Square {
	return Square{
		start: Point{
			x: xx,
			y: yy,
		},
		a: a,
	}

}

func (a Square) End() Point {
	end := Point{}
	end.x = a.start.x + int(a.a)
	end.y = a.start.y + int(a.a)
	return end
}

func (a Square) Area() uint {
	return a.a * a.a
}

func (a Square) Perimeter() uint {
	return a.a * 4
}

package main

/* Для заданной структуры квадрата создайте методы для возврата
 конечной точки, площади и периметра, используя заданные сигнатуры.
ПРИМЕЧАНИЕ: приёмник является заполнителем, замените
 соответствующим образом */

import (
	"fmt"

	square "golang-united-school-homework-5.1"
)

func main() {

	b := square.NewSquare(312, 534, 10)

	fmt.Println(b)
	fmt.Println(b.End())
	fmt.Println(b.Area())
	fmt.Println(b.Perimeter())
}

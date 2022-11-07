// 24.	Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Объявляем новую структуру
type Point struct {
	// с полями
	x float32
	y float32
}

// функция инициализирует новую переменную и возвращает на нее указатель
func NewPoint(x, y float32) *Point {
	return &Point{x: x, y: y}
}

// функция рассчитывает растояние
func Distance(p1, p2 *Point) (distance float64) {

	dx := p2.x - p1.x
	dy := p2.y - p1.y

	distance = math.Sqrt(
		math.Pow(float64(dx), 2) +
			math.Pow(float64(dy), 2))
	return
}

// функция ввода информации от пользователя
func Scan() (f1 float32, f2 float32, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	var result [2]float32
	var txt string
	for scanner.Scan() {
		txt = scanner.Text()
		if len(txt) > 0 {
			values := strings.Split(txt, ",")
			if len(values) > 2 || len(values) < 2 {
				return 0, 0, errors.New("need two values")
			} else {
				for i, v := range values {
					fl, err := strconv.ParseFloat(strings.Trim(v, " "), 32)
					if err != nil {
						return 0, 0, errors.New("value not converted to number")
					}
					result[i] = float32(fl)
				}
			}
			return result[0], result[1], nil
		}
	}
	return 0, 0, errors.New("no data")
}

func main() {
	var err error

	fmt.Print("Input coordinates for Point 1 (x,y): ")
	x1, y1, err := Scan()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input coordinates for Point 2 (x,y): ")
	x2, y2, err := Scan()
	if err != nil {
		log.Fatal(err)
	}

	a := NewPoint(x1, y1)
	b := NewPoint(x2, y2)

	fmt.Printf("Distance between Point 1 and Point 2: ")
	fmt.Println(Distance(a, b))

}

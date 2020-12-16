package main

import "fmt"

//черепаха, которая может двигаться вверх, вниз, налево или направо. Черепаха должна сохранить координаты (x, y) места
type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}

//положительные значения для передвижения вниз и направо
func (t *turtle) down() {
	t.y++
}

func (t *turtle) left() {
	t.x--
}

//положительные значения для передвижения вниз и направо
func (t *turtle) right() {
	t.x++
}

//черепаха, которая может двигаться вверх, вниз, налево или направо. имеет координаты (x, y) места
func main() {
	var t turtle
	t.up()
	t.up()
	t.left()
	t.left()
	fmt.Println(t) // Выводит окончательное местоположение: {-2 -2}
	t.down()
	t.down()
	t.right()
	t.right()
	fmt.Println(t) // Выводит окончательное местоположение: {0 0}
	t.down()
	t.right()
	t.up()
	t.up()
	fmt.Println(t) // Выводит окончательное местоположение: {1 -1}
}

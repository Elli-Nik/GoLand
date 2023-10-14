/*
Описать дом
1. Величины
2. Мебель 5 объектов
3. Бытовая техника 5 объектов
4. Состав семьи
5. Одна функция вызова, которой можно увидеть все
6. Выложить в гитхаб
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	house1 := House{
		ceilingHeight: 260,
		square:        45,
		numRooms:      3,
		furniture:     []string{"Диван", "Комод", "Кровать", "Кресло", "Стол"},
		tech:          []string{"Стиральная машина", "Холодильник", "Микроволновая печь", "Телевизор"},
		family:        []string{"Я", "Кошка", "Валентина"},
	}
	house1.showHouse()
}

type House struct {
	family        []string
	furniture     []string
	tech          []string
	ceilingHeight float64
	square        float64
	numRooms      int
}

func (h House) showHouse() {
	fmt.Println("Площадь кв. = ", h.square, "кв.м")
	fmt.Println("Высота потолка = ", h.ceilingHeight, "см")
	fmt.Println("Количество комнат = ", h.numRooms)
	fmt.Println("Мебель в доме: ", strings.Trim(fmt.Sprint(h.furniture), "[]\n"))
	fmt.Println("Техника: ", strings.Trim(fmt.Sprint(h.tech), "[]\n"))
	fmt.Println("Жильцы: ", strings.Trim(fmt.Sprint(h.family), "[]"))
}

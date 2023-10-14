package main

import (
	"fmt"
	"strings"
)

func main() {
	house1 := house{
		ceilingHeight: 260,
		square:        45,
		numRooms:      3,
		furniture:     []string{"Диван", "Комод", "Кровать", "Кресло", "Стол"},
		tech:          []string{"Стиральная машина", "Холодильник", "Микроволновая печь", "Телевизор"},
		family:        []string{"Я", "Кошка", "Валентина"},
	}
	house1.showHouse()
}

type house struct {
	family        []string
	furniture     []string
	tech          []string
	ceilingHeight float64
	square        float64
	numRooms      int
}

func (h house) showHouse() {
	fmt.Println("Площадь кв. = ", h.square, "кв.м")
	fmt.Println("Высота потолка = ", h.ceilingHeight, "см")
	fmt.Println("Количество комнат = ", h.numRooms)
	fmt.Println("Мебель в доме: ", strings.Trim(fmt.Sprint(h.furniture), "[]\n"))
	fmt.Println("Техника: ", strings.Trim(fmt.Sprint(h.tech), "[]\n"))
	fmt.Println("Жильцы: ", strings.Trim(fmt.Sprint(h.family), "[]"))

}
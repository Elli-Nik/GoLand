/*Описать дом
1. Величины
2. Мебель 5 объектов
3. Бытовая техника 5 объектов
4. Состав семьи
5. Одна функция вызова, которой можно увидеть все
6. Выложить в гитхаб
*/

package main

import "fmt"

func main() {
	house1 := house{
		ceilingHeight: 260,
		square:        45,
		numRooms:      3,
		furniture:     []string{"Диван", "Комод", "Кровать", "Кресло", "Стол"},
		tech:          []string{"Стиральная машина", "Холодильник", "Микроволновая печь", "Телевизор"},
		family:        []string{"Я", "Кошка"},
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
	fmt.Println("Количество комнат = ", h.numRooms, "\n")

	fmt.Println("Мебель в доме: ")
	for index, furniture := range h.furniture {
		fmt.Println(index+1, furniture)
	}
	fmt.Println("\n")

	fmt.Print("Бытовая техника: \n")
	for index, tech := range h.tech {
		fmt.Println(index+1, tech)
	}
	fmt.Println("\n")

	fmt.Println("Жильцы: ")
	for index, family := range h.family {
		fmt.Println(index+1, family)
	}
}

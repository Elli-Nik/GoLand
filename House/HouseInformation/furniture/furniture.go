package furniture

import "fmt"

type Sofa struct {
	name     string
	color    string
	softness string
	length   int
	width    int
}

type Chair struct {
	name  string
	color string
}

type Table struct {
	name     string
	color    string
	material string
}

type Wardrobe struct {
	name     string
	color    string
	material string
}

type Lamp struct {
	name string
}

func FillSofa() []Sofa {
	sofa1 := Sofa{"Диван", "оранжевый", "умеренно жесткий", 220, 180}
	sofa2 := Sofa{"Диван", "зеленый", "умеренно жесктий", 220, 180}
	sofa3 := Sofa{"Диван", "зеленый", "умеренно жесктий", 150, 80}
	return []Sofa{sofa1, sofa2, sofa3}
}

func FillChair() []Chair {
	chair1 := Chair{"Кресло", "зеленое"}
	return []Chair{chair1}
}

func FillTable() []Table {
	table1 := Table{"Письменный стол", "коричневый", "дерево"}
	table2 := Table{"Обеденный стол", "коричневый", "дерево"}
	return []Table{table1, table2}
}

func FillWardrobe() []Wardrobe {
	wardrobe1 := Wardrobe{"Шкаф", "коричневый", "дерево"}
	wardrobe2 := Wardrobe{"Шкаф", "рыжевато-коричневый", "дерево"}
	return []Wardrobe{wardrobe1, wardrobe2}
}

func FillLamp() []Lamp {
	lamp1 := Lamp{"Люстра"}
	lamp2 := Lamp{"Настольная лампа"}
	lamp3 := Lamp{"Торшер"}
	return []Lamp{lamp1, lamp2, lamp3}
}

func ShowSofa(Sofa []Sofa) {
	fmt.Printf("На чем мы спим:\n")
	for _, sofa := range Sofa {
		fmt.Printf("Название - %s\n Цвет - %s\n  Уровень мягкости - %s\n  Длина - %d\n  Ширина - %d\n ", sofa.name, sofa.color, sofa.softness, sofa.length, sofa.width)
	}
	fmt.Print("\n")
}

func ShowChair(Chair []Chair) {
	fmt.Printf("На чем мы сидим:\n")
	for _, chair := range Chair {
		fmt.Printf("Название - %s\n Цвет - %s\n", chair.name, chair.color)
	}
	fmt.Print("\n")
}

func ShowTable(Table []Table) {
	fmt.Printf("На чем мы пишем:\n")
	for _, table := range Table {
		fmt.Printf("Название - %s\n Цвет - %s\n  Материал - %s\n", table.name, table.color, table.material)
	}
	fmt.Print("\n")
}

func ShowWardrobe(Wardrobe []Wardrobe) {
	fmt.Printf("Где храним одежду:\n")
	for _, wardrobe := range Wardrobe {
		fmt.Printf("Название - %s\n Цвет - %s\n  Материал - %s\\n\"", wardrobe.name, wardrobe.color, wardrobe.material)
	}
	fmt.Print("\n")
}

func ShowLamp(Lamp []Lamp) {
	fmt.Printf("Источники света:\n")
	for _, lamp := range Lamp {
		fmt.Printf("Название - %s\n", lamp.name)
	}
	fmt.Print("\n")
}

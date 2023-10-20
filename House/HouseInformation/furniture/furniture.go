package furniture

import "fmt"

type Furniture struct {
	name string
}

func FillFurniture() []Furniture {
	sofa := Furniture{
		name: "Диван",
	}
	chestOfDrawers := Furniture{
		name: "Комод",
	}
	bed := Furniture{
		name: "Кровать",
	}
	armchair := Furniture{
		name: "Кресло",
	}
	desk := Furniture{
		name: "Стол",
	}
	return []Furniture{sofa, chestOfDrawers, bed, armchair, desk}
}

func ShowFurniture(Furniture []Furniture) {
	fmt.Printf("Мебель в доме:\n")
	for _, furniture := range Furniture {
		fmt.Printf(" Название - %s\n", furniture.name)
	}
	fmt.Print("\n")
}

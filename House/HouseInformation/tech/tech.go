package tech

import "fmt"

type Tech struct {
	name      string
	color     string
	model     string
	power     int
	powerName string
}

func FillTech() []Tech {
	washingMachine := Tech{
		name:  "Стиральная машина",
		model: "Zanussi",
		color: "Белый",
	}
	fridge := Tech{
		name:  "Холодильник",
		model: "LG",
		color: "Серый",
	}
	microwave := Tech{
		name:      "Микроволновая печь",
		model:     "Rolsen",
		power:     700,
		powerName: "вт",
		color:     "Серебристый",
	}
	tv := Tech{
		name:  "Телевизор",
		model: "LG",
		color: "Серый",
	}
	lamp := Tech{
		name:  "Лампа",
		color: "Белый",
	}
	return []Tech{washingMachine, fridge, microwave, tv, lamp}
}

func ShowTech(Tech []Tech) {
	fmt.Printf("Техника в доме:\n")
	for _, tech := range Tech {
		fmt.Printf("- %s\n", tech.name, tech.model, tech.power, tech.powerName, tech.color)
	}
	fmt.Print("\n")
}

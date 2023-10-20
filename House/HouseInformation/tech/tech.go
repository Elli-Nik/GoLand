package tech

import "fmt"

type Tech struct {
	name string
}

func FillTech() []Tech {
	washingMachine := Tech{
		name: "Стиральная машина",
	}
	fridge := Tech{
		name: "Холодильник",
	}
	microwave := Tech{
		name: "Микроволновая печь",
	}
	tv := Tech{
		name: "Телевизор",
	}
	lamp := Tech{
		name: "Лампа",
	}
	return []Tech{washingMachine, fridge, microwave, tv, lamp}
}

func ShowTech(Tech []Tech) {
	fmt.Printf("Техника в доме:\n")
	for _, tech := range Tech {
		fmt.Printf("- %s\n", tech.name)
	}
	fmt.Print("\n")
}

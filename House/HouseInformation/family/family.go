package family

import "fmt"

type Family struct {
	name          string
	age           int
	qualification string
	hobby         string
}

func FillFamily() []Family {
	member1 := Family{"Лиза", 20, "Студент", "tt"}
	member2 := Family{"Валентина", 76, "Пенсионер", "рисование"}
	cat := Family{"Фея", 2, "кошка", "фанатично спать"}
	return []Family{member1, member2, cat}
}

func ShowFamily(Family []Family) {
	fmt.Printf("Кто обитает в доме:\n")
	for _, members := range Family {
		fmt.Printf("Имя - %s\n Возраст - %d\n Кем является - %s\n Хобби - %s\n", members.name, members.age, members.qualification, members.hobby)
	}
	fmt.Print("\n")
}

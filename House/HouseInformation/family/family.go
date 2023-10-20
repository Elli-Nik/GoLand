package family

import "fmt"

type Family struct {
	name          string
	age           int
	qualification string
}

func FillFamily() []Family {
	member1 := Family{
		name:          "Лиза",
		age:           20,
		qualification: "Студент",
	}
	member2 := Family{
		name:          "Валентина",
		age:           76,
		qualification: "Пенсионер",
	}
	member3 := Family{
		name:          "Фея",
		age:           2,
		qualification: "Кошка",
	}
	return []Family{member1, member2, member3}
}

func ShowFamily(Family []Family) {
	fmt.Printf("Жильцы в доме:\n")
	for _, members := range Family {
		fmt.Printf("Имя - %s\n Возраст - %d\n Кем является - %s\n", members.name, members.age, members.qualification)
	}
	fmt.Print("\n")
}

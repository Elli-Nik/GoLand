package roomsAndsizes

import (
	"fmt"
)

type RoomsSizes struct {
	name  string
	value int
}

func FillRoomsSizes() []RoomsSizes {
	ceiling := RoomsSizes{
		name:  "Высота потолков, см: ",
		value: 260,
	}
	square := RoomsSizes{
		name:  "Площадь, кв.м: ",
		value: 50,
	}
	numRooms := RoomsSizes{
		name:  "Количество комнат: ",
		value: 3,
	}
	return []RoomsSizes{ceiling, square, numRooms}
}

func ShowRoomsSizes(RoomsSizes []RoomsSizes) {
	for _, RoomsSizes1 := range RoomsSizes {
		fmt.Printf("%s\n %i\n", RoomsSizes1.name, RoomsSizes1.value)
	}
}

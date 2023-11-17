package roomsAndsizes

import (
	"fmt"
)

type RoomsSizes struct {
	name  string
	value int
}

func FillRoomsSizes() []RoomsSizes {
	ceiling := RoomsSizes{"Высота потолков, см: ", 260}
	square := RoomsSizes{"Площадь, кв.м: ", 50}
	numRooms := RoomsSizes{name: "Количество комнат: ", value: 3}

	return []RoomsSizes{ceiling, square, numRooms}
}

func ShowRoomsSizes(RoomsSizes []RoomsSizes) {
	for _, roomsSizes1 := range RoomsSizes {
		fmt.Printf("%s\n %d\n", roomsSizes1.name, roomsSizes1.value)
	}
	fmt.Print("\n")
}

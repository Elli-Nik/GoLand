package fullHouse

import (
	"House/HouseInformation/family"
	"House/HouseInformation/furniture"
	"House/HouseInformation/roomsAndsizes"
	"House/HouseInformation/tech"
)

func FullHouse() {
	Family := family.FillFamily()
	Furniture := furniture.FillFurniture()
	Tech := tech.FillTech()
	RoomsSizes := roomsAndsizes.FillRoomsSizes()

	family.ShowFamily(Family)
	furniture.ShowFurniture(Furniture)
	tech.ShowTech(Tech)
	roomsAndsizes.ShowRoomsSizes(RoomsSizes)

}

package fullHouse

import (
	"House/HouseInformation/family"
	"House/HouseInformation/furniture"
	"House/HouseInformation/roomsAndsizes"
	"House/HouseInformation/tech"
)

func FullHouse() {
	Family := family.FillFamily()
	Tech := tech.FillTech()
	RoomsSizes := roomsAndsizes.FillRoomsSizes()
	Sofa := furniture.FillSofa()
	Chair := furniture.FillChair()
	Table := furniture.FillTable()
	Wardrobe := furniture.FillWardrobe()
	Lamp := furniture.FillLamp()

	family.ShowFamily(Family)
	tech.ShowTech(Tech)
	roomsAndsizes.ShowRoomsSizes(RoomsSizes)
	furniture.ShowSofa(Sofa)
	furniture.ShowChair(Chair)
	furniture.ShowTable(Table)
	furniture.ShowWardrobe(Wardrobe)
	furniture.ShowLamp(Lamp)
}

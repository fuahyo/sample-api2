package models

type AvailableRoom struct {
	RoomId           int    `db:"roomId" json:"roomId"`
	RoomName         string `db:"roomName" json:"roomName"`
	RoomCapacity     int    `db:"roomCapacity" json:"roomCapacity"`
	RoomLocationName string `db:"roomLocationName" json:"roomLocationName"`
	RoomTypeName     string `db:"roomTypeName" json:"roomTypeName"`
	RoomZoneName     string `db:"roomZoneName" json:"roomZoneName"`
	RoomStatus       string `db:"roomStatus" json:"roomStatus"`
	RoomEventColor   string `db:"roomEventColor" json:"roomEventColor"`
}

package models

type Reservation struct {
	DetailResId           int    `db:"detailResId" json:"detailResId" `
	HeaderResId           int    `db:"headerResId" json:"headerResId" `
	DetailStartDatetime   string `db:"detailStartDatetime" json:"detailStartDatetime" `
	DetailEndDatetime     string `db:"detailEndDatetime" json:"detailEndDatetime" `
	Purpose               string `db:"purpose" json:"purpose" `
	RoomId                int    `db:"roomId" json:"roomId" `
	ReservationStatus     string `db:"reservationStatus" json:"reservationStatus"`
	RecordCreatedById     int    `db:"recordCreatedById" json:"recordCreatedById" `
	RequestedRoomCapacity int    `db:"requestedRoomCapacity" json:"requestedRoomCapacity" `
	QrCode                string `db:"qrCode" json:"qrCode"`
}

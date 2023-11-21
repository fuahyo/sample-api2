// migration.go
package migration

import (
	"log"
	"sample-api2/database"
	"sample-api2/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func Migrate() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// db.AutoMigrate(&models.AvailableRoom{})

	// data := models.AvailableRoom{}
	// if db.Find(&data).RecordNotFound() {
	// 	fmt.Println("=================== run seeder user ======================")
	// 	// seederUser()
	// }
}

func seederRoom(db *gorm.DB) {
	data := models.AvailableRoom{
		RoomName:         "MR 3202",
		RoomCapacity:     4,
		RoomLocationName: "35 Floor",
		RoomTypeName:     "Jakarta",
		RoomZoneName:     "0123456789",
		RoomStatus:       "0123456789",
		RoomEventColor:   "0123456789",
	}

	db.Create(&data)
}

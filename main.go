package main

import (
	"log"
	"sample-api2/database"
	"sample-api2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	reservationHandler := handlers.ReservationHandler{DB: db}
	roomHandler := handlers.RoomHandler{DB: db}

	r := gin.Default()
	r.GET("/dashboard", reservationHandler.GetAllReservations)

	v1 := r.Group("v1")
	v1.GET("/dashboardz", reservationHandler.GetAllReservationsz)
	v1.GET("/getRooms", roomHandler.GetAvailableRoomsByReservationDetail)
	v1.GET("/dashboard2", reservationHandler.GetAllReservations2)
	v1.POST("/reservations", reservationHandler.CreateReservation)
	v1.PUT("/reservations/:id", reservationHandler.UpdateReservation)
	v1.DELETE("/reservations/:id", reservationHandler.DeleteReservation)

	if err := r.Run(":1234"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

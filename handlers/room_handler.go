package handlers

import (
	"net/http"
	"sample-api2/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type RoomHandler struct {
	DB *sqlx.DB
	// validate *validator.Validate
}

type ApiResponseData struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data []models.AvailableRoom `json:"data"`
}

func (h *RoomHandler) GetAvailableRoomsByReservationDetail(c *gin.Context) {
	var apiResponse ApiResponseData

	var availableRoom []models.AvailableRoom
	var reservation models.Reservation
	if err := c.BindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Perform a single validation on the detailStartDatetime field
	// if err := h.validate.Var(reservation.DetailStartDatetime, "required"); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errorrr": err.Error()})
	// 	return
	// }

	query := `
		EXEC spGetAvailableRoomsByReservationDetail 
		@detailStartDatetime = ?, 
		@detailEndDatetime = ?, 
		@requestedRoomCapacity = ?
	`

	if err := h.DB.Select(&availableRoom, query,
		reservation.DetailStartDatetime,
		reservation.DetailEndDatetime,
		reservation.RequestedRoomCapacity,
	); err != nil {
		apiResponse.Status.Code = http.StatusInternalServerError
		apiResponse.Status.Message = "Failed to fetch rooms"
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse.Status.Code = http.StatusOK
	apiResponse.Status.Message = "success"
	apiResponse.Data = availableRoom

	c.JSON(http.StatusOK, apiResponse)
}

// func NewRoomHandler(db *sqlx.DB) *RoomHandler {
// 	return &RoomHandler{
// 		DB:       db,
// 		validate: validator.New(),
// 	}
// }

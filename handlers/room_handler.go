package handlers

import (
	"net/http"
	"sample-api2/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gopkg.in/go-playground/validator.v9"
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

// CustomValidator struct to hold the validator
type CustomValidator struct {
	validator *validator.Validate
}

// ValidateStruct validates a struct and returns error if validation fails
func (cv *CustomValidator) ValidateStruct(obj interface{}) error {
	return cv.validator.Struct(obj)
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (h *RoomHandler) GetAvailableRoomsByReservationDetail(c *gin.Context) {
	var apiResponse ApiResponseData
	var availableRoom []models.AvailableRoom
	var reservation models.Reservation

	// Bind query parameters
	if err := c.ShouldBindQuery(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Custom validation
	cv := NewCustomValidator()
	if err := cv.ValidateStruct(reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationError": err.Error()})
		return
	}

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

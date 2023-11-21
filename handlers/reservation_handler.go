package handlers

import (
	"log"
	"net/http"
	"sample-api2/models"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type ReservationHandler struct {
	DB *sqlx.DB
	// validate *validator.Validate // Add this line
}

func (h *ReservationHandler) GetAllReservationsz(c *gin.Context) {
	var reservations []models.Reservation
	err := h.DB.Select(&reservations, "SELECT detailResId, qrCode FROM ReservationDetail")
	if err != nil {
		log.Printf("Error fetching reservations from the database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

type ApiResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		Reservation []models.Reservation `json:"reservations"`
	} `json:"data"`
}

func (h *ReservationHandler) GetAllReservations(c *gin.Context) {
	var apiResponse ApiResponse

	var reservations []models.Reservation
	err := h.DB.Select(&reservations, "SELECT detailResId, qrCode FROM ReservationDetail WHERE detailResId <= 2")
	if err != nil {
		apiResponse.Status.Code = http.StatusInternalServerError
		apiResponse.Status.Message = "Failed to fetch reservations"
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse.Status.Code = http.StatusOK
	apiResponse.Status.Message = "success"
	apiResponse.Data.Reservation = reservations

	c.JSON(http.StatusOK, apiResponse)
}

type ApiResponse2 struct {
	Data []models.Reservation `json:"reservations"`
}

func (h *ReservationHandler) GetAllReservations2(c *gin.Context) {
	var apiResponse2 ApiResponse2

	var reservations []models.Reservation
	err := h.DB.Select(&reservations, "SELECT detailResId, qrCode FROM ReservationDetail WHERE detailResId <= 4")
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiResponse2)
		return
	}
	apiResponse2.Data = reservations

	c.JSON(http.StatusOK, apiResponse2)
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {

	var reservation models.Reservation
	err := c.BindJSON(&reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if reservation.DetailStartDatetime == "" || reservation.Purpose == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required fields are missing"})
		return
	}

	_, err = h.DB.Exec("INSERT INTO ReservationDetail (headerResId, detailStartDatetime, detailEndDatetime, purpose, roomId, reservationStatus, recordCreatedById) VALUES (?, ?, ?, ?, ?, ?, ?)",
		reservation.HeaderResId,
		reservation.DetailStartDatetime,
		reservation.DetailEndDatetime,
		reservation.Purpose,
		reservation.RoomId,
		reservation.ReservationStatus,
		reservation.RecordCreatedById)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	// c.JSON(http.StatusCreated, reservation)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"value":  reservation,
	})
}

func (h *ReservationHandler) UpdateReservation(c *gin.Context) {
	id := c.Param("id")

	var reservation models.Reservation
	err := c.BindJSON(&reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate the Purpose field
	// if err := h.validate.Var(&reservation.Purpose, "required"); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error param": err.Error()})
	// 	return
	// }

	// Update the database
	_, err = h.DB.Exec("UPDATE ReservationDetail SET purpose=? WHERE detailResId=?", reservation.Purpose, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated success",
	})
}

func (h *ReservationHandler) DeleteReservation(c *gin.Context) {
	id := c.Param("id")

	_, err := h.DB.Exec("DELETE FROM ReservationDetail WHERE detailResId=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "deleted success",
		"code":   http.StatusOK,
	})
}

package handlers

import (
	"net/http"
	"sample-api2/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type PersonHandler struct {
	DB *sqlx.DB
}

func (h *PersonHandler) GetAllPersons(c *gin.Context) {
	var persons []models.Person
	err := h.DB.Select(&persons, "SELECT id, first_name, last_name FROM person")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch persons"})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err = h.DB.Exec("INSERT INTO person (first_name, last_name) VALUES (?, ?)", person.FirstName, person.LastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")

	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err = h.DB.Exec("UPDATE person SET first_name=?, last_name=? WHERE id=?", person.FirstName, person.LastName, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update person"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	_, err := h.DB.Exec("DELETE FROM person WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete person"})
		return
	}

	c.Status(http.StatusNoContent)
}

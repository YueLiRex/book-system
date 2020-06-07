package routes

import (
	"net/http"

	"book-system/models"

	"github.com/gin-gonic/gin"
)

//MakeReservation a controller to receive reservation json message and save to database
func MakeReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.NamedExec(`INSERT INTO reservation (startDatetime, duration, location) VALUES (:startDatetime, :duration, :location)`, reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, reservation)
}

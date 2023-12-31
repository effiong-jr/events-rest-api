package routes

import (
	"net/http"
	"strconv"

	"example.com/events-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEventsHandler(c *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Events fetched successfully", "events": events})
}

func getEventHandler(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event. Please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event fetched successfully", "event": event})
}

func createEventHandler(c *gin.Context) {

	var newEvent models.Event

	err := c.ShouldBindJSON(&newEvent)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "An error occured", "error": err})
		return
	}

	userId := c.GetInt64("userId")

	newEvent.UserID = userId

	err = newEvent.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event. Try again later.", "error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": newEvent})

}

// Handle delete event

func deleteEventHandler(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event

	userId := c.GetInt64("userId")

	if event.UserID != userId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized to delete this event"})
		return
	}

	err = event.DeleteEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})

}

func updateEventHandler(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	var event models.Event

	if c.GetInt64("userId") != event.UserID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event"})
		return
	}

	if err = c.ShouldBindJSON(&event); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong.", "error": err.Error()})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong. Please try again"})
		return
	}

	err = event.UpdateEventHandler()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update event failed. Please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})

}

package controller

import (
	"dareAPI/model"
	"dareAPI/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error

type DareHandler struct {
	*repositories.DareRepo
}

// CreateDareHandler is a HandlerFunc that takes user request and create valid data to the database
func (d *DareHandler) CreateDareHandler(c *gin.Context) {
	var dare model.Dare
	if err = c.ShouldBindJSON(&dare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = d.CreateDare(&dare); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dare)
}

// GetAllDaresHandler is a HandlerFunc that gets all the dares in the database
func (d *DareHandler) GetAllDaresHandler(c *gin.Context) {
	dares, err := d.GetAllDares()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dares)
}

// GetDareHandler is a HandlerFunc that gets the dare from the database according to the given id
func (d *DareHandler) GetDareHandler(c *gin.Context) {
	var dare *model.Dare
	id := c.Param("id")
	dare, err = d.GetDareByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, *dare)

}

// UpdateDareHandler is a HandlerFunc that updates the dare from the database according to the given id
func (d *DareHandler) UpdateDareHandler(c *gin.Context) {
	var updatedDare *model.Dare
	id := c.Param("id")

	if err = c.ShouldBindJSON(&updatedDare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = d.UpdateDare(id, updatedDare.Question)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, *updatedDare)
}

// DeleteDareHandler is a HandlerFunc that deletes the dare from the database according to the given id
func (d *DareHandler) DeleteDareHandler(c *gin.Context) {
	id := c.Param("id")
	err = d.DeleteDare(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "dare successfully deleted",
	})
}

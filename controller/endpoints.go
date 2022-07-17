package controller

import (
	"dareAPI/model"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	ErrIDInvalid = DictionaryErr("couldn't find the given ID, ID invalid")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func GinHome(c *gin.Context) {
	c.JSON(http.StatusOK, model.Message{Message: "welcome to drunk dares"})
}

func GetAllDares(c *gin.Context) {
	allDares := make(model.DareContainer, 0)
	c.JSON(http.StatusOK, allDares)
}

func GetDareByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	allDares := make(model.DareContainer, 0)

	for _, dare := range allDares {
		if int(dare.ID) == id {
			c.JSON(http.StatusOK, dare)
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": ErrIDInvalid,
	})
}

func GetRandomDare(c *gin.Context) {
	allDares := make(model.DareContainer, 0)

	containerLength := len(allDares)
	randomPosition := rand.Intn(containerLength + 1)

	c.JSON(http.StatusOK,
		allDares[randomPosition],
	)
}

func CreateNewDare(c *gin.Context) {
	allDares := make(model.DareContainer, 0)
	var newDare model.Dare

	if err := c.ShouldBindJSON(&newDare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	allDares = append(allDares, newDare)
	c.JSON(http.StatusOK, allDares)
}

func UpdateDare(c *gin.Context) {
	allDares := make(model.DareContainer, 0)
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedDare model.Dare

	if err := c.ShouldBindJSON(updatedDare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for idx, dare := range allDares {
		if int(dare.ID) == id {
			allDares[idx] = updatedDare
			c.JSON(http.StatusOK, updatedDare)
		}
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": ErrIDInvalid,
	})
}

func DeleteDare(c *gin.Context) {
	allDares := make(model.DareContainer, 0)
	id, _ := strconv.Atoi(c.Param("id"))

	for idx, dare := range allDares {
		if int(dare.ID) == id {
			allDares = append(allDares[:idx], allDares[idx+1:]...)
		}
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": ErrIDInvalid,
	})
}

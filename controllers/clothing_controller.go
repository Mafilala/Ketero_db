
package controllers

import (
	"net/http"
	"strconv"

	"github.com/Mafilala/ketero/backend/schemas"
	"github.com/Mafilala/ketero/backend/services"
	"github.com/gin-gonic/gin"
)

func CreateClothing(c *gin.Context) {
	var req schemas.CreateClothingRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	name := req.Name
	newClothing, err := services.CreateNewClothing(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newClothing)
}

func UpdateClothing(c *gin.Context) {
	var req schemas.UpdateClothingRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

    idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	name := req.Name
	newClothing, err := services.UpdateClothing(c, id, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newClothing)
}



func GetClothingByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	clothing, err := services.GetClothingByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clothing)
}

func DeleteClothing(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	deletedID, err := services.DeleteClothing(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_id": deletedID})
}

func GetAllClothing(c *gin.Context) {
	clothingList, err := services.GetAllClothing(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clothingList)
}

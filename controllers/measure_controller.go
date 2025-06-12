
package controllers

import (
    "net/http"
    "github.com/Mafilala/ketero/backend/services"
    "github.com/Mafilala/ketero/backend/schemas"
    "github.com/gin-gonic/gin"
    "strconv"
)

func CreateMeasure(c *gin.Context) {
	var req schemas.CreateMeasureRequest
	if err := c.BindJSON(&req); err != nil {
	       c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	    return
	}
	
	name := req.Name

	new_measure, err := services.CreateNewMeasure(c, name)	

	if err != nil {
	    c.JSON(500, gin.H{"error": err.Error()}) 
	    return
	}

	c.JSON(http.StatusCreated, new_measure) 
}

func DeleteMeasure(c *gin.Context) {
	idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	measure_id, err := services.DeleteMeasure(c, id)

	if err != nil {
	    c.JSON(500, gin.H{"error": "Failed to delete measure"})
	    return
	}

	c.JSON(http.StatusOK, measure_id)
}

func GetMeasureByID(c *gin.Context) {
	idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

        measure, err := services.GetMeasureByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, measure)
}

func GetAllMeasures(c *gin.Context) {
        measures, err := services.GetAllMeasures(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, measures)
}

func UpdateMeasure(c *gin.Context) {
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
	updatedMeasure, err := services.UpdateMeasure(c, id, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, updatedMeasure)
}



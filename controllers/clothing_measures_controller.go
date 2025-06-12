package controllers

import (
    "net/http"
    "github.com/Mafilala/ketero/backend/services"
    "github.com/Mafilala/ketero/backend/schemas"
    "github.com/gin-gonic/gin"
    "strconv"
)

func AddMeasure(c *gin.Context) {
	var req schemas.CreatClothingMeasureRequest
	if err := c.BindJSON(&req); err != nil {
	       c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	    return
	}
	
	clothing_id := req.Clothing_id
	measure_id := req.Measure_id

	addedMeasure, err := services.AddMeasure(c, clothing_id, measure_id)	

	if err != nil {
	    c.JSON(500, gin.H{"error": err.Error()}) 
	    return
	}

	c.JSON(http.StatusCreated, addedMeasure) 

}


func RemoveMeasure(c *gin.Context) {
    measureIdParam := c.Param("clothing_id")
	idParam := c.Param("measure_id")
    clothing_id, err1 := strconv.Atoi(measureIdParam)
    measure_id, err2 := strconv.Atoi(idParam)
 
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ids"})
		return
	}

	
	deleted_measure_id, err := services.RemoveMeasure(c, clothing_id, measure_id)

	if err != nil {
	    c.JSON(500, gin.H{"error": err})
	    return
	}

	c.JSON(http.StatusOK, deleted_measure_id)
}

func GetAllClothingMeasures(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

    all_clothing_measures, err := services.GetAllClothingMeasures(c, id)
    if err != nil {
	c.JSON(500, gin.H{"error": err})
    }

    c.JSON(http.StatusOK, all_clothing_measures)
}  

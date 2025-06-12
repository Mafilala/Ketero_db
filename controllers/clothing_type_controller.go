package controllers

import (
    "net/http"
    "github.com/Mafilala/ketero/backend/services"
    "github.com/Mafilala/ketero/backend/schemas"
    "github.com/gin-gonic/gin"
    "strconv"

)

func CreateClothingType(c *gin.Context) {
	var req schemas.CreateClothingTypeRequest
	if err := c.BindJSON(&req); err != nil {
	       c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	    return
	}
	
	name := req.Name

	new_clothing_type, err := services.CreateNewClothingType(c, name)	

	if err != nil {
	    c.JSON(500, gin.H{"error": err.Error()}) 
	    return
	}

	c.JSON(http.StatusCreated, new_clothing_type) 

}


func GetClothingTypeByID(c *gin.Context) {
	idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	clothing_type, err := services.GetClothingTypeByID(c, id)

	if err != nil {
	    c.JSON(500, gin.H{"error": err})
	    return
	}

	c.JSON(http.StatusOK, clothing_type)
}

func DeleteClothingType(c *gin.Context) {
	idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	clothing_type_id, err := services.DeleteClothingType(c, id)

	if err != nil {
	    c.JSON(500, gin.H{"error": err})
	    return
	}

	c.JSON(http.StatusOK, clothing_type_id)
}

func UpdateClothingType(c *gin.Context) {
	var req schemas.UpdateClothingTypeRequest
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
	newClothingType, err := services.UpdateClothingType(c, id, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newClothingType)
}


func GetAllClothingTypes(c *gin.Context) {
    all_clothing_types, err := services.GetAllClothingTypes(c)
    if err != nil {
	c.JSON(500, gin.H{"error": err})
    }

    c.JSON(http.StatusOK, all_clothing_types)
}

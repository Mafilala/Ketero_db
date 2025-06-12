package controllers

import (
    "net/http"
    "github.com/Mafilala/ketero/backend/services"
    "github.com/Mafilala/ketero/backend/schemas"
    "github.com/gin-gonic/gin"
    "strconv"

)

func AddClothing(c *gin.Context) {
	var req schemas.CreatClothingTypePartRequest
	if err := c.BindJSON(&req); err != nil {
	       c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	    return
	}
	
	clothing_type_id := req.Clothing_type_id
	clothing_id := req.Clothing_id

	addedClothing, err := services.AddClothing(c, clothing_type_id, clothing_id)	

	if err != nil {
	    c.JSON(500, gin.H{"error": err.Error()}) 
	    return
	}

	c.JSON(http.StatusCreated, addedClothing) 

}


func RemoveClothingTypePart(c *gin.Context) {
	partIdParam := c.Param("clothing_type_id")
	idParam := c.Param("clothing_id")
        clothing_type_id, err1 := strconv.Atoi(partIdParam)
        clothing_id, err2 := strconv.Atoi(idParam)
 
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ids"})
		return
	}

	deleted_clothing_id, err := services.RemoveClothingTypePart(c, clothing_type_id, clothing_id)

	if err != nil {
	    c.JSON(500, gin.H{"error": err})
	    return
	}

	c.JSON(http.StatusOK, deleted_clothing_id)
}

func GetAllClothingParts(c *gin.Context) {
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

    all_clothing_parts, err := services.GetAllClothingParts(c, id)
    if err != nil {
	c.JSON(500, gin.H{"error": err})
    }

    c.JSON(http.StatusOK, all_clothing_parts)
}  

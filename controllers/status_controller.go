

package controllers

import (
    "net/http"
    "strconv"

    "github.com/Mafilala/ketero/backend/schemas"
    "github.com/Mafilala/ketero/backend/services"
    "github.com/gin-gonic/gin"
)

func CreateStatus(c *gin.Context) {
    var req schemas.CreateStatusRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    newStatus, err := services.CreateNewStatus(c, req.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newStatus)
}

func GetStatusByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    status, err := services.GetStatusByID(c, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, status)
}

func DeleteStatus(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    deletedID, err := services.DeleteStatus(c, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"deleted_id": deletedID})
}

func GetAllStatuses(c *gin.Context) {
    statuses, err := services.GetAllStatuses(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, statuses)
}

func UpdateStatus(c *gin.Context) {
	var req schemas.UpdateStatus
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
	updatedStatus, err := services.UpdateStatus(c, id, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, updatedStatus)
}



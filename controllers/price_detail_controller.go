
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Mafilala/ketero/backend/schemas"
	"github.com/Mafilala/ketero/backend/services"
)

func CreatePriceDetail(c *gin.Context) {
	var req schemas.CreatePriceDetailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	detail, err := services.CreatePriceDetail(c, req.ToModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create price detail"})
		return
	}

	c.JSON(http.StatusCreated, detail)
}

func GetPriceDetailByOrderID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	detail, err := services.GetPriceDetailByOrderID(c, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

func UpdatePriceDetail(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	var req schemas.UpdatePriceDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Fetch current record to preserve unspecified fields
	existing, err := services.GetPriceDetailByOrderID(c, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Price detail not found"})
		return
	}

	updated := req.ToModel(orderID, *existing)

	result, err := services.UpdatePriceDetail(c, updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update price detail"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeletePriceDetail(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	err = services.DeletePriceDetail(c, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete price detail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Price detail deleted successfully"})
}

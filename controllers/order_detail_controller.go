
package controllers

import (
	"net/http"
	"strconv"

	"github.com/Mafilala/ketero/backend/schemas"
	"github.com/Mafilala/ketero/backend/models"
	"github.com/Mafilala/ketero/backend/services"
	"github.com/gin-gonic/gin"
)

func CreateOrderDetail(c *gin.Context) {
	var req schemas.CreateOrderDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	detail := models.OrderDetail{
		OrderID: req.OrderID,
		Style:   req.Style,
		Fabric:  req.Fabric,
		Color:   req.Color,
	}

	result, err := services.CreateOrderDetail(c.Request.Context(), detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order detail"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateOrderDetail(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	var req schemas.UpdateOrderDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	current, err := services.GetOrderDetail(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order detail not found"})
		return
	}

	if req.Style != nil {
		current.Style = req.Style
	}
	if req.Fabric != nil {
		current.Fabric = req.Fabric
	}
	if req.Color != nil {
		current.Color = req.Color
	}

	if err := services.UpdateOrderDetail(c.Request.Context(), *current); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order detail"})
		return
	}

	c.JSON(http.StatusOK, current)
}

func DeleteOrderDetail(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	if err := services.DeleteOrderDetail(c.Request.Context(), orderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order detail"})
		return
	}

	c.Status(http.StatusNoContent)
}

func GetOrderDetail(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	detail, err := services.GetOrderDetail(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order detail not found"})
		return
	}

	c.JSON(http.StatusOK, detail)
}

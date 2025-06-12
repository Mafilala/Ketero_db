package controllers

import (
	"net/http"
	"strconv"
	"log"
	"github.com/Mafilala/ketero/backend/schemas"
	"github.com/Mafilala/ketero/backend/services"
	"github.com/gin-gonic/gin"
)

func CreateOrderMeasure(c *gin.Context) {
	var req []schemas.CreateOrderMeasureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("create order", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	
	for _, measure := range req {
		_, err := services.CreateOrderMeasure(c, measure.ToModel())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		}

	}

	
	c.JSON(http.StatusCreated, gin.H{"message": "Measures created successfully"})}

func UpdateOrderMeasure(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	
	orderID, err1 := strconv.Atoi(orderIDStr)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req schemas.UpdateOrderMeasureRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateOrderMeasure(c, orderID, req.ClothingID, req.Measures); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func DeleteOrderMeasure(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	measureIDStr := c.Param("measure_id")

	orderID, err1 := strconv.Atoi(orderIDStr)
	measureID, err2 := strconv.Atoi(measureIDStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := services.DeleteOrderMeasure(c, orderID, measureID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func GetOrderMeasuresByOrderID(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	measures, err := services.GetOrderMeasuresByOrderID(c, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, measures)
}


package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/Mafilala/ketero/backend/models"
    "github.com/Mafilala/ketero/backend/schemas"
    "github.com/Mafilala/ketero/backend/services"
    "log"
)

func CreateOrder(c *gin.Context) {
    var req schemas.CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Println("Error inserting order:", err)

        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    order := models.Order{
        ClientID:       req.ClientID,
        ClothingTypeID: req.ClothingTypeID,
        StatusID:       req.StatusID,
        OrderNote:      req.OrderNote,
        DueDate:        req.DueDate,
    }

    newOrder, err := services.CreateNewOrder(c, order)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newOrder)
}

func GetOrderByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    order, err := services.GetOrderByID(c, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    deletedID, err := services.DeleteOrder(c, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"deleted": deletedID})
}

func GetAllOrders(c *gin.Context) {
    limitStr := c.DefaultQuery("limit", "10")
    offsetStr := c.DefaultQuery("offset", "0")
    statusStr := c.DefaultQuery("status", "")

    limit, _ := strconv.Atoi(limitStr)
    offset, _ := strconv.Atoi(offsetStr)

    orders, total, err := services.GetAllOrders(c, limit, offset, statusStr)
    if err != nil {
        log.Println("Error getting all orders:", err)

        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  orders,
        "total": total,
    })
}

func PatchOrder(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req schemas.PatchOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Println("Error inserting order:", err)

        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    order := models.Order{
        ClientID:       req.ClientID,
        ClothingTypeID: req.ClothingTypeID,
        StatusID:       req.StatusID,
        OrderNote:      req.OrderNote,
        DueDate:        req.DueDate,
    }


    updatedOrder, err := services.PatchOrder(c, id, &order)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedOrder)
}



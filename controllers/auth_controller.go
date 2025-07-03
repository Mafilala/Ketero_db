package controllers

import (
	"fmt"
	"net/http"

	"github.com/Mafilala/ketero/backend/utils"
	"github.com/gin-gonic/gin"
)

func WebAppHandler(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to read body"})
		fmt.Println("reading", err)
		return
	}

	params := string(body)
	verified := utils.VerifyTelegramWebApp(params)
	if verified {
		c.JSON(200, gin.H{"verified": true})
		return
	}

	// access granted
	c.JSON(200, gin.H{"verified": false})
}

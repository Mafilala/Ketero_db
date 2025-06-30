package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Mafilala/ketero/backend/utils"
	"github.com/gin-gonic/gin"
)

func WebAppHandler(c *gin.Context) {
	var params map[string]string
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		fmt.Println("binding", err)
		return
	}

	// Convert map to url.Values
	q := url.Values{}
	for k, v := range params {
		q.Set(k, v)
	}

	err := utils.VerifyTelegramWebApp(q)
	if err != nil {
		c.JSON(200, gin.H{"verified": false})
		return
	}

	// access granted
	c.JSON(200, gin.H{"verified": true})
}

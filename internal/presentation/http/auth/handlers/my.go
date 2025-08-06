package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/app/utils"
)

func MyHandler(c *gin.Context) {
	clientIDVal, exists := c.Get(utils.GinContextClientIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	clientID, ok := clientIDVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid clientId"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"clientId": clientID})
}

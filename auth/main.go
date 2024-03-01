package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func checkTokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "no token provided",
		})
		ctx.Abort()
		return
	}

	tokenParts := strings.Split(token, " ")
	if len(tokenParts) != 2 || tokenParts[1] != "123" {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token provided",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}

func main() {
	router := gin.Default()

	router.Use(checkTokenMiddleware)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "auth success, hello from auth service",
		})
	})

	router.Run("0.0.0.0:3000")
}

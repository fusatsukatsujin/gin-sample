package controllers

import (
	"database/sql"
	"gin-sample/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context, db *sql.DB) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	authService := services.NewAuthService(db)
	token, err := authService.Authenticate(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

package controllers

import (
	"database/sql"
	"gin-sample/pkg/setting"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(c *gin.Context, db *sql.DB) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// TODO:ここでユーザー認証を行う（例：データベースでチェック）
	if username == "admin" && password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": 1,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(setting.AppSetting.JwtSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗しました"})
	}

}

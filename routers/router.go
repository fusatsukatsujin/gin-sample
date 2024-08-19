package routers

import (
	"database/sql"
	member_controller "gin-sample/controllers"
	"net/http"

	"gin-sample/middleware"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

func InitRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "こんにちは、Gin!",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// ここでユーザー認証を行う（例：データベースでチェック）
		if username == "admin" && password == "password" {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": 1,
				"exp":     time.Now().Add(time.Hour * 24).Unix(),
			})

			tokenString, err := token.SignedString([]byte("your_secret_key"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗しました"})
		}
	})

	protected := r.Group("/api")
	// TODO: 本番環境では環境変数から取得する
	protected.Use(middleware.AuthMiddleware("your_secret_key"))

	protected.POST("/members", func(c *gin.Context) {
		member_controller.AddMember(c, db)
	})

	protected.GET("/members", func(c *gin.Context) {
		member_controller.GetMembers(c, db)
	})

	protected.GET("/members/:id", func(c *gin.Context) {
		member_controller.GetMemberById(c, db)
	})
	return r
}

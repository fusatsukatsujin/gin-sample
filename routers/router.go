package routers

import (
	"database/sql"
	"gin-sample/controllers"
	"gin-sample/pkg/setting"
	"net/http"

	"gin-sample/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	setupPublicRoutes(r, db)
	setupProtectedRoutes(r, db)

	return r
}

func setupPublicRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/", homeHandler)
	r.POST("/login", func(c *gin.Context) {
		controllers.Authenticate(c, db)
	})
}

func setupProtectedRoutes(r *gin.Engine, db *sql.DB) {
	protected := r.Group("/api")
	// TODO: 本番環境では環境変数から取得する
	protected.Use(middleware.AuthMiddleware(setting.AppSetting.JwtSecret))

	protected.POST("/members", func(c *gin.Context) {
		controllers.AddMember(c, db)
	})
	protected.GET("/members", func(c *gin.Context) {
		controllers.GetMembers(c, db)
	})
	protected.GET("/members/:id", func(c *gin.Context) {
		controllers.GetMemberById(c, db)
	})
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "こんにちは、Gin!",
	})
}

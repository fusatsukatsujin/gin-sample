package routers

import (
	"database/sql"
	"gin-sample/pkg/setting"
	controllers "gin-sample/routers/api/v1"
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
	protected.Use(middleware.AuthMiddleware(setting.AppSetting.JwtSecret))

	// トランザクション管理が必要なエンドポイントのグループ
	txGroup := protected.Group("")
	tm := middleware.NewTransactionManager(db)
	txGroup.Use(tm.HandleTransaction())

	txGroup.POST("/members", func(c *gin.Context) {
		controllers.AddMember(c, db)
	})

	// トランザクション不要なエンドポイント
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

package routes

import (
	"final/controller"
	"final/middlewares"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

func Setuprouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controller.Login)
	r.POST("register", controller.Register)

	UserMiddlewareRoute := r.Group("/user")
	UserMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	UserMiddlewareRoute.GET("", controller.Getalluser)
	UserMiddlewareRoute.GET("/:id", controller.GetuserById)
	UserMiddlewareRoute.PATCH("/:username", controller.UpdatePasswordUser)
	UserMiddlewareRoute.DELETE("/:id", controller.Deleteuser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

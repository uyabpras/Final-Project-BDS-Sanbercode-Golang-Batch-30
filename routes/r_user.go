package routes

import (
	"final/controller"

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

	r.GET("/user", controller.Getalluser)
	r.POST("/user", controller.CreateUser)
	r.GET("/user/:id", controller.GetuserById)
	r.PATCH("/user/:username", controller.UpdatePasswordUser)
	r.DELETE("/user/:id", controller.Deleteuser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

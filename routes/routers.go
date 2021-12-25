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
	r.POST("/register", controller.Register)
	r.POST("/login-store", controller.Loginstore)
	r.POST("/register-store", controller.Registerstore)
	r.POST("/update-store", controller.UpdatePasswordUser)

	UserMiddlewareRoute := r.Group("/user")
	UserMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	UserMiddlewareRoute.GET("", controller.Getalluser)
	UserMiddlewareRoute.GET("/:id", controller.GetuserById)
	UserMiddlewareRoute.PATCH("", controller.UpdatePasswordUser)
	UserMiddlewareRoute.DELETE("/:id", controller.Deleteuser)

	ProductMiddlewareRoute := r.Group("/product")
	ProductMiddlewareRoute.GET("", controller.Getallproduct)
	ProductMiddlewareRoute.GET("/price/asc", controller.GetsortpriceASC)
	ProductMiddlewareRoute.GET("/price/dsc", controller.GetsortpriceDSC)
	ProductMiddlewareRoute.GET("/:id", controller.GetproductById)
	ProductMiddlewareRoute.GET("/newest", controller.GetNewestproduct)
	ProductMiddlewareRoute.GET("/latest", controller.GetLatestproduct)
	ProductMiddlewareRoute.GET("/filters", controller.Getfiltersprice)
	//
	ProductMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ProductMiddlewareRoute.POST("", controller.CreateProduct)
	ProductMiddlewareRoute.DELETE("/:name", controller.DeleteProduct)

	StoreMiddlewareRoute := r.Group("/store")
	StoreMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	StoreMiddlewareRoute.GET("", controller.Getallstore)
	StoreMiddlewareRoute.DELETE("", controller.Deletestore)
	StoreMiddlewareRoute.GET("/newest", controller.GetNeweststore)
	StoreMiddlewareRoute.PATCH("", controller.UpdateDataStore)

	InventoryMiddlewareRoute := r.Group("/inventory")
	InventoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	InventoryMiddlewareRoute.GET("", controller.Getinventory)
	InventoryMiddlewareRoute.POST("", controller.Createinventory)
	InventoryMiddlewareRoute.GET("/asc", controller.Getinventoryasc)
	InventoryMiddlewareRoute.GET("/desc", controller.Getinventorydesc)

	AddresMiddlewareRoute := r.Group("/addres")
	AddresMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	AddresMiddlewareRoute.GET("", controller.GetallAddresbyUserID)
	AddresMiddlewareRoute.POST("", controller.CreateAddres)
	AddresMiddlewareRoute.DELETE("/:id", controller.DeleteAddresUser)
	AddresMiddlewareRoute.PATCH("/:id", controller.UpdateAddresdUser)

	OrderMiddlewareRoute := r.Group("/order")
	OrderMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	OrderMiddlewareRoute.GET("", controller.Createorder)

	CheckoutMiddlewareRoute := r.Group("/checkout")
	CheckoutMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	CheckoutMiddlewareRoute.POST("", controller.Createcheckout)
	CheckoutMiddlewareRoute.GET("", controller.Getcheckout)
	CheckoutMiddlewareRoute.GET("/asc", controller.Getcheckoutasc)
	CheckoutMiddlewareRoute.GET("/desc", controller.Getcheckoutdesc)

	CartMiddlewareRoute := r.Group("/cart")
	CartMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	CartMiddlewareRoute.POST("", controller.CreateCart)
	CartMiddlewareRoute.GET("", controller.GetallcartUser)
	CartMiddlewareRoute.DELETE("", controller.DeletecartUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

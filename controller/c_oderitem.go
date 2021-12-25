package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderInput struct {
	ProductID int `json:"productid"`
	Price     int `json:"price"`
	AddresID  int `json:"addresid"`
}

// creating  order godoc
// @Summary creating  data order
// @Description creating data order which input productID, price n addres id
// @Tags order
// @param Body body OrderInput true "the body to create new order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Orderitem
// @router /order [post]
func Createorder(ctx *gin.Context) {
	var input OrderInput

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	orders := models.Orderitem{UserID: ID, ProductID: input.ProductID, Price: input.Price, Addres_id: input.AddresID, CreatedAt: time.Now()}

	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&orders)

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

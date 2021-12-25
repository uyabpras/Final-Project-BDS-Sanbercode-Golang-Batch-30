package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartInput struct {
	UserID    int `json:"userid"`
	ProductID int `json:"productid"`
	StoreID   int `json:"storeid"`
}

// creating cart product godoc
// @Summary creating cart product
// @Description creating cart which input productid n storeid
// @Tags Cart
// @param Body body CartInput true "the body to create new cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Cart
// @router /cart [post]
func CreateCart(ctx *gin.Context) {
	var input CartInput
	db := ctx.MustGet("db").(*gorm.DB)

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	keranjang := models.Cart{UserID: userID, ProductID: input.ProductID, StoreID: input.StoreID}

	db.Create(&keranjang)

	ctx.JSON(http.StatusOK, gin.H{"data": keranjang})
}

// Get all cart by userid godoc
// @Summary geting data cart by extracting jwt token
// @Description getting data cart by extracting jwt token
// @Tags Cart
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Cart
// @router /cart [get]
func GetallcartUser(ctx *gin.Context) {
	var keranjang models.Cart
	db := ctx.MustGet("db").(*gorm.DB)

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("user_id = ?", userID).Find(&keranjang).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
}

// delete cart by userid n id cart godoc
// @Summary deleting data cart user by extracting jwt token and selecting cart for delete using cart id
// @Description deleting data cart user by extracting jwt token and selecting cart for delete using cart id
// @Tags Cart
// @Produce json
// @param id path string true "data cart by id "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Cart
// @router /cart/:id [delete]
func DeletecartUser(ctx *gin.Context) {
	var keranjang models.Cart
	db := ctx.MustGet("db").(*gorm.DB)

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("user_id = ? and id = ?", userID, ctx.Param("id")).Find(&keranjang).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	db.Delete(&keranjang)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

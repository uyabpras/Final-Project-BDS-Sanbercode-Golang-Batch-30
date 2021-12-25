package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CheckoutInput struct {
	StoreID int `json:"storeid"`
	OrderID int `json:"orderid"`
}

// creating  checkout godoc
// @Summary creating  data checkout
// @Description creating data Checkout which input storeID n orderID
// @Tags Checkout
// @param Body body CheckoutInput true "the body to create new Checkout"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Checkout
// @router /checkout [post]
func Createcheckout(ctx *gin.Context) {
	var input CheckoutInput

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	checkouts := models.Checkout{StoreID: ID, OrderitemID: input.OrderID, CreatedAt: time.Now()}

	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&checkouts)

	ctx.JSON(http.StatusOK, gin.H{"data": checkouts})
}

// Get data checkout by id store godoc
// @Summary geting data checlout by extracting jwt token
// @Description getting data checkout by extracting jwt token
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Checkout
// @router /checkout [get]
func Getcheckout(ctx *gin.Context) {
	var checkouts models.Checkout
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Find(&checkouts).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": checkouts})
}

// Get data checkout by id store(ascending) godoc
// @Summary geting data checlout by extracting jwt token (ascending)
// @Description getting data checkout by extracting jwt token (ascending)
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Checkout
// @router /checkout/asc [get]
func Getcheckoutasc(ctx *gin.Context) {
	var checkouts models.Checkout
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Order("creared_at asc").Find(&checkouts).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": checkouts})
}

// Get data checkout by id store(descending) godoc
// @Summary geting data checlout by extracting jwt token (descending)
// @Description getting data checkout by extracting jwt token (descending)
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Checkout
// @router /checkout/desc [get]
func Getcheckoutdesc(ctx *gin.Context) {
	var checkouts models.Checkout
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Order("creared_at desc").Find(&checkouts).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": checkouts})
}

package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreInput struct {
	Name string `json:"name"`
	City string `json:"city"`
}

// Get All data store
// @Summary geting all data store
// @Description getting all data store
// @Tags store
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Store
// @router /store [get]
func Getallstore(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var toko []models.Store

	db.Find(&toko)

	ctx.JSON(http.StatusOK, gin.H{"data": toko})
}

// Get storename by name godoc
// @Summary delete data store by name
// @Description delete data store by name
// @Tags store
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @router /store [delete]
func Deletestore(ctx *gin.Context) {
	var toko models.Store
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Find(&toko).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.Delete(&toko)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

// Get sort newest store
// @Summary Get newest data store
// @Description Get newest data store
// @Tags store
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Store
// @router /store/newest [get]
func GetNeweststore(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var toko []models.Store

	db.Order("created_at desc").Find(&toko)
	ctx.JSON(http.StatusOK, gin.H{"data": toko})
}

// Update store
// @Summary update data store
// @Description update store
// @Tags Store
// @Produce json
// @param Body body StoreInput true "the body to updated data store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.User
// @router /store [patch]
func UpdateDataStore(ctx *gin.Context) {
	// check data
	var keranjang models.Store
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	if err := db.Where("id = ?", ID).Find(&keranjang).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	//input data
	var input StoreInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var updatedata models.Store

	updatedata.Username = input.Name
	updatedata.City = input.City
	updatedata.UpdateAt = time.Now()

	db.Model(&keranjang).Updates(updatedata)
	ctx.JSON(http.StatusOK, gin.H{"data": keranjang})
}

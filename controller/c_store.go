package controller

import (
	"final/models"
	"net/http"

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
// @Success 200 {object} []models.Store
// @router /store [get]
func Getallstore(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var toko []models.Store

	db.Find(&toko)

	ctx.JSON(http.StatusOK, gin.H{"data": toko})
}

// creating  store godoc
// @Summary creating  data store
// @Description creating data store which input store name n city
// @Tags store
// @param Body body StoreInput true "the body to create new store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Store
// @router /store [post]
func Createstore(ctx *gin.Context) {
	var input StoreInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	stores := models.Store{Name: input.Name, City: input.City}
	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&stores)

	ctx.JSON(http.StatusOK, gin.H{"data": stores})
}

// Get storename by name godoc
// @Summary delete data store by name
// @Description delete data store by name
// @Tags store
// @Produce json
// @param id path string true "data store by name "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @router /store/{name} [delete]
func Deletestore(ctx *gin.Context) {
	var toko models.Store
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("name = ?", ctx.Param("name")).Find(&toko).Error; err != nil {
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
// @Success 200 {object} []models.Store
// @router /store/newest [get]
func GetNeweststore(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var toko []models.Store

	db.Order("created_at desc").Find(&toko)
	ctx.JSON(http.StatusOK, gin.H{"data": toko})
}

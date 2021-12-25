package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventoryInput struct {
	ProuctID int `json:"productid"`
	Stock    int `json:"stock"`
}

// Get data inventory by id store godoc
// @Summary geting data checlout by extracting jwt token
// @Description getting data inventory by extracting jwt token
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Inventory
// @router /inventory [get]
func Getinventory(ctx *gin.Context) {
	var inventory models.Inventory
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Find(&inventory).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": inventory})
}

// creating  inventory godoc
// @Summary creating  data inventory
// @Description creating data inventory which input storeID n productID
// @Tags inventory
// @param Body body InventoryInput true "the body to create new inventory"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Inventory
// @router /inventory [post]
func Createinventory(ctx *gin.Context) {
	var input InventoryInput

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	inventorys := models.Inventory{StoreID: ID, ProductID: input.ProuctID, CreatedAt: time.Now()}

	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&inventorys)

	ctx.JSON(http.StatusOK, gin.H{"data": inventorys})
}

// Get data checkout by id store(descending) godoc
// @Summary geting data checlout by extracting jwt token (descending)
// @Description getting data inventory by extracting jwt token (descending)
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Inventory
// @router /inventory/desc [get]
func Getinventoryasc(ctx *gin.Context) {
	var inventorys models.Inventory
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Order("creared_at asc").Find(&inventorys).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": inventorys})
}

// Get data checkout by id store(descending) godoc
// @Summary geting data checlout by extracting jwt token (descending)
// @Description getting data inventory by extracting jwt token (descending)
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Inventory
// @router /inventory/desc [get]
func Getinventorydesc(ctx *gin.Context) {
	var inventorys models.Inventory
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Order("creared_at desc").Find(&inventorys).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": inventorys})
}

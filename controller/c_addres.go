package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddresInput struct {
	UserId     int    `json:"userid"`
	Addresline string `json:"addresline"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
}

// Get all addres by userid godoc
// @Summary geting data addres by extracting jwt token
// @Description getting data addres by extracting jwt token
// @Tags Addres
// @Produce json
// @param id path string true "data address by id "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Addres
// @router /addres [get]
func GetallAddresbyUserID(ctx *gin.Context) {
	var alamat models.Addres
	db := ctx.MustGet("db").(*gorm.DB)

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("user_id = ?", userID).Find(&alamat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

}

// creating  Addres godoc
// @Summary creating  data Addres
// @Description creating data Addres for order item which includes addressline, city, state and country
// @Tags Addres
// @param Body body AddresInput true "the body to create new Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Addres
// @router /addres [post]
func CreateAddres(ctx *gin.Context) {
	var input AddresInput

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	alamat := models.Addres{UserID: userID, Addresline: input.Addresline, City: input.City,
		State: input.State, Country: input.Country}
	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&alamat)

	ctx.JSON(http.StatusOK, gin.H{"data": alamat})
}

// delete addres by userid n id addres godoc
// @Summary deleting data addres user by extracting jwt token and selecting addres for delete using addres id
// @Description deleting data addres user by extracting jwt token and selecting addres for delete using addres id
// @Tags Addres
// @Produce json
// @param id path string true "data address by id "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Addres
// @router /addres/:id [delete]
func DeleteAddresUser(ctx *gin.Context) {
	var alamat models.Addres
	db := ctx.MustGet("db").(*gorm.DB)

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("user_id = ? and id = ?", userID, ctx.Param("id")).Find(&alamat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	db.Delete(&alamat)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

// Update addres
// @Summary update addres user by extracting jwt token
// @Description update addres user by extracting jwt token
// @Tags User
// @Produce json
// @param username path string true "update addres"
// @param Body body AddresInput true "the body to updated addres"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Addres
// @router /addres/:id [patch]
func UpdateAddresdUser(ctx *gin.Context) {
	// check data
	var alamat models.Addres
	db := ctx.MustGet("db").(*gorm.DB)

	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("user_id = ? and id = ?", userID, ctx.Param("id")).Find(&alamat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	//input data
	var input AddresInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var updatedata models.Addres

	updatedata.UserID = userID
	updatedata.Addresline = input.Addresline
	updatedata.City = input.City
	updatedata.State = input.State
	updatedata.Country = input.Country
	updatedata.UpdateAt = time.Now()

	db.Model(&alamat).Updates(updatedata)
	ctx.JSON(http.StatusOK, gin.H{"data": alamat})
}

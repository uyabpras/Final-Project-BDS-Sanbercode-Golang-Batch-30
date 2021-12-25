package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Get All User godoc
// @Summary geting all data user
// @Description getting all data user which includes username n password
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.User
// @router /User [get]
func Getalluser(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var users []models.User

	db.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Get Username by id godoc
// @Summary geting data user by extracting jwt token
// @Description getting data user by extracting jwt token which includes username n password
// @Tags User
// @Produce json
// @param id path string true "data user by id "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.User
// @router /User [get]
func GetuserById(ctx *gin.Context) {
	var users models.User
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Find(&users).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Get Username by extracting jwt token godoc
// @Summary delete data user by extracting jwt token
// @Description delete data user by extracting jwt token which includes username n password
// @Tags User
// @Produce json
// @param id path string true "data user by extracting jwt token "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @router /User [delete]
func Deleteuser(ctx *gin.Context) {
	var users models.User
	db := ctx.MustGet("db").(*gorm.DB)
	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Find(&users).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.Delete(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

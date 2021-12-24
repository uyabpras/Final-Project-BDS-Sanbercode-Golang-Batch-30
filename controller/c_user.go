package controller

import (
	"final/models"
	"final/utils/token"
	"net/http"
	"time"

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

// creating All User godoc
// @Summary creating all data user
// @Description creating all data user which input username n password
// @Tags User
// @param Body body UserInput true "the body to create new user"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.User
// @router /User [post]
func CreateUser(ctx *gin.Context) {
	var input UserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users := models.User{Username: input.Username, Password: input.Password}
	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&users)

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
// @router /User/ [get]
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

// Update passowrd
// @Summary update data by extracting jwt token
// @Description update password user by extracting jwt token
// @Tags User
// @Produce json
// @param username path string true "update password by username"
// @param Body body UserInput true "the body to updated password"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.User
// @router /User [patch]
func UpdatePasswordUser(ctx *gin.Context) {
	// check data
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
	//input data
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var updatedata models.User

	updatedata.Password = input.Password
	updatedata.UpdateAt = time.Now()

	db.Model(&users).Updates(updatedata)
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
// @router /User/{id} [delete]
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

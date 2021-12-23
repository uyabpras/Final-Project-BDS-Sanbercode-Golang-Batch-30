package controller

import (
	"final/models"
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
// @Produce json
// @Success 200 {object} []models.User
// @router /User [post]
func CreateUser(ctx *gin.Context) {
	var input UserInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	users := models.User{Username: input.Username, Password: input.Password}
	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Get Username by id godoc
// @Summary geting data user by id
// @Description getting data user by id which includes username n password
// @Tags User
// @Produce json
// @param id path string true "data user by id "
// @Success 200 {object} []models.User
// @router /User/{id} [get]
func GetuserById(ctx *gin.Context) {
	var users models.User
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", ctx.Param("id")).Find(&users).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Update passowrd
// @Summary update data by Username
// @Description update password user by username
// @Tags User
// @Produce json
// @param username path string true "update password by username"
// @param Body body UserInput true "the body to updated password"
// @Success 200 {object} []models.User
// @router /User/username [patch]
func UpdatePasswordUser(ctx *gin.Context) {
	// check data
	var users models.User
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("user_username = ?", ctx.Param("username")).Find(&users).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//input data
	var input UserInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var updatedata models.User

	updatedata.Username = input.Username
	updatedata.Password = input.Password
	updatedata.UpdateAt = time.Now()

	db.Model(&users).Updates(updatedata)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Get Username by id godoc
// @Summary delete data user by id
// @Description delete data user by id which includes username n password
// @Tags User
// @Produce json
// @param id path string true "data user by id "
// @Success 200 {object} map[string]boolean
// @router /User/{id} [delete]
func Deleteuser(ctx *gin.Context) {
	var users models.User
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", ctx.Param("id")).Find(&users).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.Delete(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

package controller

import (
	"final/models"
	"final/utils/token"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserAuth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StoreAuth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	City     string `json:"city" binding:"required"`
}

type ChangePassword struct {
	Password string `json:"password" binding:"required"`
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body UserAuth true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var input UserAuth

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	user := map[string]string{
		"username": u.Username,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body UserAuth true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input UserAuth

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// update password user godoc
// @Summary update password user.
// @Description update a user from public access.
// @Tags Auth
// @Param Body body ChangePassword true "the body to update password user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user [patch]
func UpdatePasswordUser(ctx *gin.Context) {
	var input ChangePassword
	var u models.User
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := db.Where("id = ?", ID).Find(&u).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Password = input.Password
	u.UpdateAt = time.Now()

	_, err = u.UpdateUser(db)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"status": "succes changed password",
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "update success", "user": user})

}

// Loginstore godoc
// @Summary Login as as store.
// @Description Logging in to get jwt token to access admin or store api by roles.
// @Tags Auth
// @Param Body body StoreAuth true "the body to login a store"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login-store [post]
func Loginstore(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var input UserAuth

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Store{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheckstore(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	user := map[string]string{
		"username": u.Username,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}

// Register store godoc
// @Summary Register a store.
// @Description registering a store from public access.
// @Tags Auth
// @Param Body body StoreAuth true "the body to register a store"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register-store [post]
func Registerstore(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input StoreAuth

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Store{}

	u.Username = input.Username
	u.Password = input.Password
	u.City = input.City
	u.CreatedAt = time.Now()

	_, err := u.SaveUserstore(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// update password store godoc
// @Summary update password store.
// @Description update a store from public access.
// @Tags Auth
// @Param Body body ChangePassword true "the body to update password store"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /update-store [patch]
func UpdatePasswordstore(ctx *gin.Context) {
	var input ChangePassword
	var u = models.Store{}
	db := ctx.MustGet("db").(*gorm.DB)

	ID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	if err := db.Where("id = ?", ID).Find(&u).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Password = input.Password
	u.UpdateAt = time.Now()

	_, err = u.UpdateUserstore(db)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"status": "succes changed password",
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "update success", "user": user})

}

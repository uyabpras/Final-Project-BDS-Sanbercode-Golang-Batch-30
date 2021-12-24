package controller

import (
	"final/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Get All data product
// @Summary geting all data product
// @Description getting all data product
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @router /product [get]
func Getallproduct(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var produk []models.Product

	db.Find(&produk)

	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// Get Product by id godoc
// @Summary geting data Product by id
// @Description getting data Product by id
// @Tags Product
// @Produce json
// @param id path string true "data Product by id "
// @Success 200 {object} []models.Product
// @router /product/{id} [get]
func GetproductById(ctx *gin.Context) {
	var produk []models.Product
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", ctx.Param("id")).Find(&produk).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// Get sort price product by asscending
// @Summary Get sort price product by asscending
// @Description Get sort price product by asscending
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @router /product/price/asc [get]
func GetsortpriceASC(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var produk []models.Product

	db.Order("price asc").Find(&produk)
	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// Get sort price product by descending
// @Summary Get sort price product by descending
// @Description Get sort price product by descending
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @router /product/price/dsc [get]
func GetsortpriceDSC(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var produk []models.Product

	db.Order("price desc").Find(&produk)
	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// Get sort newest product
// @Summary Get newest data product
// @Description Get newest data product
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @router /product/newest [get]
func GetNewestproduct(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var produk []models.Product

	db.Order("created_at desc").Find(&produk)
	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// Get sort latest product
// @Summary Get latest data product
// @Description Get latest data product
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @router /product/latest [get]
func GetLatestproduct(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var produk []models.Product

	db.Order("created_at asc").Find(&produk)
	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// Get filters price product
// @Summary Get filters price product
// @Description Get filters price product, using param= "price1" for ">=" and "price2" for "<="
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @router /product/filters [get]
func Getfiltersprice(ctx *gin.Context) {
	var produk []models.Product
	db := ctx.MustGet("db").(*gorm.DB)
	price1 := ctx.Query("price1")
	price2 := ctx.Query("price2")

	db.Where("price >= ? and price <= ?", price1, price2).Find(&produk)

	ctx.JSON(http.StatusOK, gin.H{"data": produk})
}

// creating  Product godoc
// @Summary creating  data Product
// @Description creating data Product which input Product name n price
// @Tags Product
// @param Body body ProductInput true "the body to create new Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Product
// @router /product [post]
func CreateProduct(ctx *gin.Context) {
	var input ProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	Products := models.Product{Name: input.Name, Price: input.Price}
	db := ctx.MustGet("db").(*gorm.DB)

	db.Create(&Products)

	ctx.JSON(http.StatusOK, gin.H{"data": Products})
}

// Get Productname by id godoc
// @Summary delete data Product by id
// @Description delete data Product by id which includes Productname n password
// @Tags Product
// @Produce json
// @param id path string true "data Product by name "
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @router /product/{name} [delete]
func DeleteProduct(ctx *gin.Context) {
	var Products models.Product
	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("name = ?", ctx.Param("name")).Find(&Products).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.Delete(&Products)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}

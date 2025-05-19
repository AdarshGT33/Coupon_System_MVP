// @title Coupon System API
// @version 1.0
// @description This is the backend for validating and applying coupons.
// @termsOfService http://yourdomain.com/terms/

// @contact.name Adarsh Singh Tomar
// @contact.email tomaradarsh18@gmail.com

// @host localhost:8080
// @BasePath /
package main

import (
	"net/http"

	_ "coupon_system/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"coupon_system/cache"
	"coupon_system/controller"
	"coupon_system/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
	utils.ConnectDB()
}

func main() {
	r := gin.Default()

	cache.InitCouponCache()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "we are up!",
		})
	})
	r.POST("/create-coupon", controller.CreateCoupon)
	r.POST("/validate-coupon", controller.ValidateCoupon)

	r.Run()
}

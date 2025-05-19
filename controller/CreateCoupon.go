package controller

import (
	"net/http"

	"coupon_system/helper"
	"coupon_system/models"
	"coupon_system/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCouponHandler godoc
// @Summary Create a new coupon
// @Description Allows creation of a new coupon with full configuration
// @Tags Coupon
// @Accept json
// @Produce json
// @Param coupon body models.Coupon true "Coupon payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create-coupon [post]
func CreateCoupon(c *gin.Context) {
	var input helper.CreateCouponInput

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "Invalid input",
			"Details": err.Error(),
		})
		return
	}

	var medicines []models.Medicine
	if len(input.ApplicableMedicineIDs) > 0 {
		err := utils.DB.Where("id IN ?", input.ApplicableMedicineIDs).Find(&medicines).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch medicines"})
			return
		}
	}

	coupon := models.Coupon{
		ID:                   uuid.New(),
		CouponCode:           input.CouponCode,
		ExpiryDate:           input.ExpiryDate,
		UsageType:            input.UsageType,
		ApplicableCategories: input.ApplicableCategories,
		MinOrderValue:        input.MinOrderValue,
		ValidFrom:            input.ValidTimeWindow.Start,
		ValidTo:              input.ValidTimeWindow.End,
		TermsAndCondition:    input.TermsAndConditions,
		DiscountType:         input.DiscountType,
		DiscountValue:        input.DiscountValue,
		MaxUsagePerUser:      &input.MaxUsagePerUser,
		ApplicableMedicines:  medicines,
	}

	if err := utils.DB.Create(&coupon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating coupon",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "coupon created",
		"coupon":  coupon,
	})
}

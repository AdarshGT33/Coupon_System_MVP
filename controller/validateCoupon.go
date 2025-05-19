package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gocache "github.com/patrickmn/go-cache"

	"coupon_system/cache"
	"coupon_system/helper"
	"coupon_system/models"
	"coupon_system/utils"
)

// ValidateCoupon godoc
// @Summary Validate a coupon
// @Description Validates a coupon for a specific order and checks all conditions
// @Tags Coupon
// @Accept  json
// @Produce  json
// @Param coupon body helper.UserSentCoupon true "Coupon input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /validate-coupon [post]
func ValidateCoupon(c *gin.Context) {
	var user_coupon helper.UserSentCoupon
	err := c.ShouldBindBodyWithJSON(&user_coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coupon"})
		return
	}

	//TTL caching to reduce DB calls
	var coupon *models.Coupon
	if cachedCoupon, found := cache.CouponCache.Get(user_coupon.CouponCode); found {
		coupon = cachedCoupon.(*models.Coupon)
	} else {
		var dbCoupon models.Coupon
		error := utils.DB.Preload("ApplicableMedicines").Where("coupon_code = ?", user_coupon.CouponCode).First(&dbCoupon).Error
		if error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
			return
		}
		cache.CouponCache.Set(user_coupon.CouponCode, &dbCoupon, gocache.DefaultExpiration)
		coupon = &dbCoupon
	}

	// Lock coupon usage based on coupon ID
	couponLock := helper.CouponLocks(coupon.ID.String())
	couponLock.Lock()
	defer couponLock.Unlock()

	// Expiry check
	if time.Now().After(coupon.ExpiryDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Coupon Expired!"})
		return
	}

	// Validity window check
	if time.Now().Before(coupon.ValidFrom) || time.Now().After(coupon.ValidTo) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Coupon not valid at this time"})
		return
	}

	// Minimum order amount check
	if user_coupon.OrderValue < coupon.MinOrderValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart does not reach the minimum amount"})
		return
	}

	// Medicine applicability check
	couponMedicineMap := make(map[uuid.UUID]bool)
	for _, med := range coupon.ApplicableMedicines {
		couponMedicineMap[med.ID] = true
	}

	matchFound := false
	for _, medID := range user_coupon.MedicineIDs {
		if couponMedicineMap[medID] {
			matchFound = true
			break
		}
	}

	if !matchFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Coupon is not applicable to any medicines in the order"})
		return
	}

	// Usage type check
	var usageCount int64
	utils.DB.Model(&models.CouponUsage{}).Where("user_id = ? AND coupon_id = ?", user_coupon.UserID, coupon.ID).Count(&usageCount)

	switch coupon.UsageType {
	case "one_time":
		if usageCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Coupon can only be used one time"})
			return
		}
	case "multi_use":
		if coupon.MaxUsagePerUser != nil && usageCount >= int64(*coupon.MaxUsagePerUser) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Max usage limit reached for this coupon"})
			return
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid coupon usage type"})
		return
	}

	// Discount calculation
	var discount float64
	switch coupon.DiscountType {
	case "flat":
		discount = coupon.DiscountValue
	case "percentage":
		discount = (coupon.DiscountValue / 100) * user_coupon.OrderValue
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid discount type"})
		return
	}

	finalAmount := user_coupon.OrderValue - discount
	if finalAmount < 0 {
		finalAmount = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Coupon applied successfully",
		"coupon_code":    coupon.CouponCode,
		"discount":       discount,
		"final_amount":   finalAmount,
		"original_price": user_coupon.OrderValue,
	})
}

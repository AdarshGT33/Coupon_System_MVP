package helper

import (
	"time"

	"coupon_system/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ValidTimeWindow struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type CreateCouponInput struct {
	CouponCode            string              `json:"coupon_code"`
	ExpiryDate            time.Time           `json:"expiry_date"`
	UsageType             models.UsageType    `json:"usage_type"`
	ApplicableMedicineIDs []uuid.UUID         `json:"applicable_medicine_ids"`
	ApplicableCategories  pq.StringArray      `json:"applicable_categories" gorm:"type:text[]"`
	MinOrderValue         float64             `json:"min_order_value"`
	ValidTimeWindow       ValidTimeWindow     `json:"valid_time_window"`
	TermsAndConditions    string              `json:"terms_and_conditions"`
	DiscountType          models.DiscountType `json:"discount_type"`
	DiscountValue         float64             `json:"discount_value"`
	MaxUsagePerUser       int                 `json:"max_usage_per_user"`
}

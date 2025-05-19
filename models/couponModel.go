package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UsageType string

const (
	OneTime   UsageType = "one_time"
	MultiUse  UsageType = "multi_use"
	TimeBased UsageType = "time_based"
)

type DiscountType string

const (
	Flat       DiscountType = "flat"
	Percentage DiscountType = "percentage"
)

type Coupon struct {
	ID                   uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CouponCode           string         `gorm:"uniqueIndex;not null" json:"coupon_code"`
	ExpiryDate           time.Time      `json:"expiry_date"`
	UsageType            UsageType      `gorm:"type:varchar(20)" json:"usage_type"`
	ApplicableMedicines  []Medicine     `gorm:"many2many:coupon_medicine" json:"applicable_medicines"`
	ApplicableCategories pq.StringArray `gorm:"type:text[]" json:"applicable_categories"`
	MinOrderValue        float64        `json:"min_order_value"`
	ValidFrom            time.Time      `json:"valid_from"`
	ValidTo              time.Time      `json:"valid_to"`
	TermsAndCondition    string         `gorm:"type:text" json:"terms_and_conditions"`
	DiscountType         DiscountType   `gorm:"type:varchar(20)" json:"discount_type"`
	DiscountValue        float64        `json:"discount_value"`
	MaxUsagePerUser      *int           `json:"max_usage_per_user"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

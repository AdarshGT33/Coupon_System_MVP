package helper

import (
	"time"

	"github.com/google/uuid"
)

type UserSentCoupon struct {
	UserID      uuid.UUID   `json:"user_id"`
	CouponCode  string      `json:"coupon_code"`
	MedicineIDs []uuid.UUID `json:"medicine_ids"`
	OrderValue  float64     `json:"order_value"`
	OrderTime   time.Time   `json:"order_time"`
}

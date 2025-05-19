package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var CouponCache *cache.Cache

func InitCouponCache() {
	CouponCache = cache.New(5*time.Minute, 10*time.Minute)
}

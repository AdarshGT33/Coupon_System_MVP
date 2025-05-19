package helper

import "sync"

var couponLocks = make(map[string]*sync.Mutex)
var couponLocksMu sync.Mutex

func CouponLocks(couponID string) *sync.Mutex {
	couponLocksMu.Lock()
	defer couponLocksMu.Unlock()

	if _, exist := couponLocks[couponID]; !exist {
		couponLocks[couponID] = &sync.Mutex{}
	}
	return couponLocks[couponID]
}

package utils

import (
	"fmt"
	"imooc_go_lottery/datasource"
)

func getLuckyLockKey(uid int) string {
	return fmt.Sprintf("lucky_lock_%d", uid)
}

func LockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	// key value exists time expire(s) 设置过期时间避免死锁
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	}
	return false
}

func UnlockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("DEL", key)
	if rs == "OK" {
		return true
	}
	return false
}


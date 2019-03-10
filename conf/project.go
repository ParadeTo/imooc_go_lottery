package conf

import "time"

const SysTimeForm = "2006-01-02 15:04:05"
const SysTimeFormShort = "2016-01-02"

const (
	GtypeVirtual = iota // 虚拟币
	GtypeCodeSame       // 虚拟券，相同的码
	GtypeCodeDiff				// 虚拟券，不同的码
	GtypeGiftSmall      // 实物小奖
	GtypeGiftLarge      // 实物大奖
)

const IpLimitMax = 500

const IpPrizeMax = 10 // 一个 IP 可以抽奖的次数

const UserPrizeMax = 3000

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

var SignSecret = []byte("0123456789abcdef")

var CookieSecret = "hellolottery"

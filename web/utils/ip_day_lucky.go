package utils

import (
	"fmt"
	"imooc_go_lottery/comm"
	"imooc_go_lottery/datasource"
	"log"
	"math"
	"time"
)

const ipFramSize = 2

func init () {
	resetGroupIpList()
}

func resetGroupIpList () {
	log.Println("ip_day_lucky.resetGroupIpList start")
	cacheObj := datasource.InstanceCache()
	for i := 0; i < ipFramSize; i++  {
		key := fmt.Sprintf("day_ips_%d", i)
		cacheObj.Do("DEL", key)
	}
	log.Println("ip_day_lucky.resetGroupIpList end")
	// IP 当天的统计数，零点的时候归零，设置定时器
	duration := comm.NextDayDuration()
	time.AfterFunc(duration, resetGroupIpList)
}

func IncrIpLuckyNum(strIp string) int64 {
	ip := comm.Ip4toInt(strIp)
	i := ip % ipFramSize
	key := fmt.Sprintf("day_ips_%d", i)
	cacheObj := datasource.InstanceCache()
	// key 表示这个 hset 的域名，ip 是键值
	rs, err := cacheObj.Do("HINCRBY", key, ip, 1)
	if err != nil {
		log.Println("IncrIpLuckyNum error=", err)
		return math.MaxInt64
	} else {
		return rs.(int64)
	}
}

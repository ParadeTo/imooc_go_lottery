package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)
//var mu sync.Mutex
// 奖品中奖概率
type Prate struct {
	Rate  int // 万分之 n 的中奖概率
	Total int // 总数量限制，0 表示无限数量
	CodeA int // 中奖概率起始编码（包含）
	CodeB int // 中奖概率终止编码（包含）
	Left  *int32 // 剩余数
}

// 奖品列表
var prizeList []string = []string{
	"一等奖，火星一日游",
	"二等奖，地球半日游",
	"三等奖，月球一日游",
	"",
}

// 奖品的中奖概率
var left = int32(1000)
var rateList []Prate = []Prate {
	{10000,1,0,9999,&left},
	//{2,2,1,2,2},
	//{5,10,3,5,10},
	//{100,1,0,9999,0},
}

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get () string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("大转盘奖品列表：<br /> %s",
		strings.Join(prizeList, "<br />\n"))
}

func (c *lotteryController) GetDebug() string {
	return fmt.Sprintf("获奖概率：%v\n", rateList)
}

func (c *lotteryController) GetPrize() string {
	// 1. 抽奖，根据随机数匹配奖品
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(10000)

	var myprize string
	var prizeRate *Prate

	for i, prize := range prizeList {
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			myprize = prize
			prizeRate = rate
			break
		}
	}
	if myprize == "" {
		myprize = "很遗憾，再来一次吧"
		return myprize
	}

	// 2. 发奖
	if prizeRate.Total == 0 {
		return myprize
	} else if *prizeRate.Left > 0 {
		//mu.Lock()
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			return myprize
		}
		//prizeRate.Left -= 1
		//mu.Unlock()
	}
	myprize = "很遗憾，再来一次吧"
	return myprize
}

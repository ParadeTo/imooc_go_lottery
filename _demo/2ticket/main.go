/**
1 即开记得型
2 双色球自选型
 */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"time"
)

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

// 即开即得
func (c *lotteryController) Get() string {
	var prize string
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)
	switch {
	case code == 1:
		prize = "一等奖"
	case code >= 2 && code <= 3:
		prize = "二等奖"
	case code >= 4 && code <= 6:
		prize = "三等奖"
	default:
		return fmt.Sprintf("尾号为1获得一等奖<br/>"+
			"尾号为2或者3获得二等奖<br/>"+
			"尾号为4/5/6获得三等奖<br/>"+
			"code=%d<br/>"+
			"很遗憾，没有获奖", code)
	}
	return fmt.Sprintf("尾号为1获得一等奖<br/>"+
		"尾号为2或者3获得二等奖<br/>"+
		"尾号为4/5/6获得三等奖<br/>"+
		"code=%d<br/>"+
		"恭喜你获得：%s", code, prize)
}

func (c *lotteryController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var prize [7]int
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1
	}
	prize[6] = r.Intn(16) + 1
	return fmt.Sprintf("今日开奖号码是: %v", prize)
}
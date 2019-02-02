package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type gift struct {
	id       int
	name     string
	pic      string
	link     string
	inuse    bool
	rate     int
	rateMin  int
	rateMax  int
}

const rateMax = 10000

var logger *log.Logger

var giftList *[5]gift
var mu sync.Mutex

type lotteryController struct {
	Ctx iris.Context
}

func initLog() {
	f, _ := os.Create("./lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newGift() *[5]gift {
	giftList = new([5]gift)
	g1 := gift{
		id:       1,
		name:     "富强福",
		pic:      "",
		link:     "",
		inuse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[0] = g1
	g2 := gift{
		id:       2,
		name:     "和谐福",
		pic:      "",
		link:     "",
		inuse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[1] = g2
	g3 := gift{
		id:       3,
		name:     "友善福",
		pic:      "",
		link:     "",
		inuse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[2] = g3
	g4 := gift{
		id:       4,
		name:     "爱国福",
		pic:      "",
		link:     "",
		inuse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[3] = g4
	g5 := gift{
		id:       5,
		name:     "敬业福",
		pic:      "",
		link:     "",
		inuse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[4] = g5
	return giftList
}

func giftRage(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)
	rateStart := 0
	for i, data := range giftList {
		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + data.rate
		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
	fmt.Printf("giftList=%v\n", giftList)
	return giftList
}

func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRage(rate)
	return fmt.Sprintf("%v\n", giftList)
}

func (c *lotteryController) GetLucky() map[string]interface{} {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	uid, _ := c.Ctx.URLParamInt("uid")
	code := luckyCode()
	ok := false
	giftList := giftRage(rate)
	result := make(map[string]interface{})
	result["success"] = ok
	for _, data := range giftList {
		if !data.inuse {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			sendData := data.name
			saveLuckyData(code, data.id, data.name, data.link, sendData)
			result["success"] = ok
			result["id"] = data.id
			result["name"] = data.name
			result["link"] = data.link
			result["uid"] = uid
			result["data"] = sendData
			break
		}
	}
	return result
}

func saveLuckyData(code int32, id int, name, link, sendData string) {
	logger.Printf("lucky, code=%d, id=%d, name=%s, link=%s, sendData=%s\n",
		code, id, name, link, sendData)
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

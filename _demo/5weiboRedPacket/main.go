package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"sync"
	"time"
)

type task struct {
	id uint32
	callback chan uint
}

var logger *log.Logger
const taskNum = 16
var packageList *sync.Map = new(sync.Map)
var chTaskList []chan task = make([] chan task, taskNum)

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	for i := 0; i < taskNum; i++ {
		chTaskList[i] = make(chan task)
		go fetchPackagelistMoney(chTaskList[i])
	}
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get () map[uint32][2]int {
	rs := make(map[uint32][2]int)
	//for id, list := range packageList {
	//	var money int
	//	for _, v := range list {
	//		money += int(v)
	//	}
	//	rs[id] = [2]int{len(list), money}
	//}
	packageList.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		list := value.([]int)
		var money int
		for _, v := range list {
			money += int(v)
		}
		rs[id] = [2]int{len(list), money}
		return true
	})
	return rs
}

// 发红包
func (c *lotteryController) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")
	if errUid != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("参数格式异常，errUid=%d, errMoney=%d, errNum=%d\n",
			errUid, errMoney, errNum)
	}
	moneyTotal := int(money * 100) // 转为分
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%f, num=%d\n",
			uid, money, num)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.55 // 随机分配的最大值
	list := make([]uint, num)
	leftMoney := moneyTotal
	leftNum := num
	for leftNum > 0 {
		if leftNum == 1 {
			// 最后一个红包
			list[num-1] = uint(leftMoney)
		}
		// 红包最小金额为一分
		if leftMoney == leftNum {
			for i := num-leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}
		// 留出至少 leftNum 分钱
		rMoney := int(float64(leftMoney - leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}
	id := r.Uint32()
	//packageList[id] = list
	packageList.Store(id, list)
	// 返回抢红包的地址
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)
}

// 抢红包
func (c *lotteryController) GetGet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	id, errId := c.Ctx.URLParamInt("id")
	if errUid != nil || errId != nil {
		return fmt.Sprintf("参数格式异常，errUid=%d, errId=%d\n",
			errUid, errId)
	}
	if uid < 1 || id < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, id=%d\n",
			uid, id)
	}

	var list []uint
	var ok bool
	var listRaw interface{}
	if listRaw, ok = packageList.Load(uint32(id)); ok {
		list = listRaw.([]uint)
	}
	if !ok || len(list) < 1 {
		return fmt.Sprintf("红包不存在, id=%d\n", id)
	}

	// 1 构造一个抢红包任务
	callback := make(chan uint)
	t := task{id: uint32(id), callback: callback}
	chTask := chTaskList[id % taskNum]
	chTask <- t
	money := <-callback
	if money <= 0 {
		return "没有抢到红包"
	}

	// 2. 分配随机数
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//i := r.Intn(len(list))
	//money := list[i]
	//if len(list) > 1 {
	//	if i == len(list) - 1 {
	//		packageList.Store(uint32(id), list[:i])
	//	} else if i == 0 {
	//		packageList.Store(uint32(id), list[1:])
	//	} else {
	//		packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
	//	}
	//} else {
	//	packageList.Delete(uint32(id))
	//}

	return fmt.Sprintf("恭喜你抢到一个红包，金额为：%d\n", money)
}

func fetchPackagelistMoney(chTask chan task) {
	for {
		t := <-chTask
		id := t.id
		listRaw, ok := packageList.Load(uint32(id))
		if ok || listRaw != nil {
			list := listRaw.([]uint)
			// 分配随机数
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(len(list))
			money := list[i]
			if len(list) > 1 {
				if i == len(list) - 1 {
					packageList.Store(uint32(id), list[:i])
				} else if i == 0 {
					packageList.Store(uint32(id), list[1:])
				} else {
					packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
				}
			} else {
				packageList.Delete(uint32(id))
			}
			t.callback <- money
		} else {
			t.callback <- 0
		}
	}
}

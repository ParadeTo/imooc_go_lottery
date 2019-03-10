package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"imooc_go_lottery/models"
	"sync"
	"time"
)

var engine *xorm.Engine

func main() {
	engine, _ = xorm.NewEngine("mysql", "root:123456@/lottery?charset=utf8")
	engine.SetMaxOpenConns(2000)
	engine.SetMaxIdleConns(1000)
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for {
				data := &models.LtGift{Id: 1}
				engine.Get(data)
				fmt.Println(data)
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
	wg.Wait()
}
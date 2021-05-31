package main

import (
	"github.com/robfig/cron/v3"
	"go-cron/gocron"
	test_task "go-cron/test-task"
	"log"
	"time"
)

func main() {
	gocron.InitCronInstance(cron.New())
	//从mysql中读取出来
	listFromMysql := []gocron.Task{
		{
			CronSpec: "@every 1s",
			Type:     "test1",
		},
		{
			CronSpec: "@every 5s",
			Type:     "test2",
		},
	}
	for _, task := range listFromMysql {
		if err := gocron.Register(task); err != nil {
			log.Fatal(err)
		}
	}
	if err := gocron.Register(gocron.Task{
		Type:        "test1",
		ExecuteFunc: test_task.ExecuteTest1,
	}); err != nil {
		log.Fatal(err)
	}
	if err := gocron.Register(gocron.Task{
		Type:        "test2",
		ExecuteFunc: test_task.ExecuteTest2,
	}); err != nil {
		log.Fatal(err)
	}
	gocron.Start()
	time.Sleep(1 * time.Hour)
}

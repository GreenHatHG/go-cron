package gocron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type Task struct {
	//robfig/cron cron.addFunc spec参数
	CronSpec    string
	Type        string
	ExecuteFunc func()
}

//key：任务实例名字符串, value：实例
var taskTypes map[string]Task
var cronInstance *cron.Cron

func InitCronInstance(cronInst *cron.Cron) {
	cronInstance = cronInst
}

//注册任务实例
func Register(t Task) error {
	if taskTypes == nil {
		taskTypes = make(map[string]Task)
	}

	task, ok := taskTypes[t.Type]
	//该情况为已经注册了数据库的任务列表，当时ExecuteInterface没有参数
	if ok {
		if t.ExecuteFunc == nil {
			return fmt.Errorf("%s已经注册，但ExecuteInterface仍参数为nil\n", t.Type)
		}
		task.ExecuteFunc = t.ExecuteFunc
	} else {
		//全新注册的列表
		taskTypes[t.Type] = t
	}

	if err := bindTaskTrigger(task); err != nil {
		return err
	}
	return nil
}

func Start() {
	cronInstance.Start()
}

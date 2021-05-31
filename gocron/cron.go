package gocron

import "fmt"

func bindTaskTrigger(t Task) error {
	if t.ExecuteFunc == nil {
		return nil
	}
	if _, err := cronInstance.AddFunc(t.CronSpec, t.ExecuteFunc); err != nil {
		return fmt.Errorf("cron.AddFunc出错：%+v\n", err)
	}
	return nil
}

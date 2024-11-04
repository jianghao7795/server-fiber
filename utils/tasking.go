package utils

import (
	"errors"
	"time"
)

//@author: wuhao
//@function: Tasking
//@description: 执行任务
//@param:
//@return: error

func Tasking(taskName string, output string, interval string) error {
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	// log.Println("taskName: ", taskName, ", output: ", output, ", duration: ", duration)
	return err
}

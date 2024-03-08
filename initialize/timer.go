package initialize

import (
	"fmt"

	"server-fiber/config"
	"server-fiber/global"
	"server-fiber/utils"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.CONFIG.Timer.Detail[i])
		}
	}
}

func Tasks() {
	record := 10
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Tasks {
			go func(detail config.Task) {
				global.Timer.AddTaskByFunc("Tasking", global.CONFIG.Timer.Spec, func() {
					if record == 10 {
						record = 0
						global.Timer.StopTask("Tasking")
					} else {
						err := utils.Tasking(detail.TaskName, detail.Output, detail.Interval)
						record++
						if err != nil {
							fmt.Println("tasking error:", err)
						}
					}
				})
			}(global.CONFIG.Timer.Tasks[i])
		}
	}
}

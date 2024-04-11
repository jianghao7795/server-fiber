package initialize

import (
	"server-fiber/config"
	"server-fiber/global"
	"server-fiber/utils"

	"go.uber.org/zap"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("UpdateGithub", global.CONFIG.Timer.Spec, func() {
					err := utils.UpdateTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						global.LOG.Error("新增Github记录错误：", zap.Error(err))
					}
				})
			}(global.CONFIG.Timer.Detail[i])
		}
	}
}

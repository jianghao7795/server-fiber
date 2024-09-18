package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"server-fiber/global"
	"server-fiber/initialize"
	"server-fiber/service/system"

	"go.uber.org/zap"
)

func RunServer() {
	var err error
	global.VIP, err = viperInit() // 初始化Viper 配置
	if err != nil {
		// global.LOG.Error("配置错误：", zap.Error(err))
		log.Println("配置错误： ", err.Error())
		os.Exit(1)
	}
	global.LOG, err = zapInit() // 初始化zap日志库
	if err != nil {
		log.Println("日志初始化错误： ", err.Error())
	}
	// global.Logger = InitLogger()   // 初始化 log 让log标准输出
	zap.ReplaceGlobals(global.LOG) // 配置部署到全局

	db, err := initialize.Gorm() // gorm连接数据库
	if err == nil {
		global.DB = db
		global.LOG.Info("Database connection success", zap.String("port", global.CONFIG.Mysql.Port))
	} else {
		global.LOG.Error("The database connection failed: " + err.Error())
		os.Exit(1)
	}
	initialize.Timer() // 定时 执行任务
	// err = utilsInit.TransInit("zh")
	// if err != nil {
	// 	global.LOG.Error("翻译错误：" + err.Error())
	// }
	if global.DB != nil {
		system.LoadAll() // 加载所有的 拉黑的jwt数据 避免盗用jwt
		// initialize.RegisterTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.LOG.Error("数据库关闭失败: " + err.Error())
			}
		}(db)
	}
	if global.CONFIG.System.UseMultipoint || global.CONFIG.System.UseRedis {
		err = initialize.Redis()
		if err != nil {
			global.LOG.Error("Redis init failed: " + err.Error())
			os.Exit(1)
		}
	}
	router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	global.LOG.Info("server run success on ", zap.String("address", address))
	log.Println(`Welcome to Fiber API`)
	global.LOG.Error(router.Listen(address).Error())
}

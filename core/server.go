package core

import (
	"fmt"

	"server-fiber/global"
	"server-fiber/initialize"
	"server-fiber/service/system"
	utilsInit "server-fiber/utils"

	"go.uber.org/zap"
)

func RunServer() {
	global.VIP = Viper() // 初始化Viper 配置
	global.LOG = Zap()   // 初始化zap日志库
	// global.Logger = core.InitLogger() // 初始化 log 让log标准输出
	zap.ReplaceGlobals(global.LOG) // 部署到全局

	db, err := initialize.Gorm() // gorm连接数据库
	if err == nil {
		global.DB = db
		global.LOG.Info("Database connection success", zap.String("port", global.CONFIG.Mysql.Port))
	} else {
		global.LOG.Error("The database connection failed: " + err.Error())
		panic(err)
	}
	initialize.Tasks() //定时 执行任务
	utilsInit.TransInit("zh")
	if global.DB != nil {
		system.LoadAll() // 加载所有的 拉黑的jwt数据 避免盗用jwt
		// initialize.RegisterTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	if global.CONFIG.System.UseMultipoint || global.CONFIG.System.UseRedis {
		initialize.Redis()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	global.LOG.Info("server run success on ", zap.String("address", address))
	fmt.Println(`Welcome to Fiber API`)
	global.LOG.Error(Router.Listen(address).Error())
}

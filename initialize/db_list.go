package initialize

import (
	"server-fiber/global"

	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.Dbname], _ = GormMysqlByConfig(info)
		case "pgsql":
			dbMap[info.Dbname] = GormPgSqlByConfig(info)
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		global.DB = sysDB
	}
	global.DBList = dbMap
}

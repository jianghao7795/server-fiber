package system

import (
	global "server-fiber/model"
	"server-fiber/model/system/response"
)

type Database interface {
	GetDB() (data []response.Db, err error)
	GetTables(dbName string) (data []response.Table, err error)
	GetColumn(tableName string, dbName string) (data []response.Column, err error)
}

func (autoCodeService *AutoCodeService) Database() Database {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return AutoCodeMysql
	case "pgsql":
		return AutoCodePgsql
	default:
		return AutoCodeMysql
	}
}

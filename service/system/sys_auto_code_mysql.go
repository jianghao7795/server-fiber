package system

import (
	global "server-fiber/model"
	"server-fiber/model/system/response"
)

// GetDB 获取数据库的所有数据库名
// Author wuhao
func (s *autoCodeMysql) GetDB() (data []response.Db, err error) {
	var entities []response.Db
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;"
	err = global.DB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author wuhao
func (s *autoCodeMysql) GetTables(dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `select table_name as table_name from information_schema.tables where table_schema = ?`
	err = global.DB.Raw(sql, dbName).Scan(&entities).Error
	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author wuhao
func (s *autoCodeMysql) GetColumn(tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := `
	SELECT COLUMN_NAME        column_name,
       DATA_TYPE          data_type,
       CASE DATA_TYPE
           WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'int' THEN c.NUMERIC_PRECISION
           WHEN 'bigint' THEN c.NUMERIC_PRECISION
           ELSE '' END AS data_type_long,
       COLUMN_COMMENT     column_comment
	FROM INFORMATION_SCHEMA.COLUMNS c
	WHERE table_name = ?
	  AND table_schema = ?
	`
	err = global.DB.Raw(sql, tableName, dbName).Scan(&entities).Error
	return entities, err
}

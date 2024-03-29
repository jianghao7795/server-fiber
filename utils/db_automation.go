package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server-fiber/global"
	"server-fiber/model/system"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

//@author: wuhao
//@function: ClearTable
//@description: 清理数据库表数据
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

func ClearTable(db *gorm.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	return db.Debug().Exec(fmt.Sprintf("update %s SET deleted_at = ? WHERE %s < ?", tableName, compareField), time.Now().Add(-duration), time.Now().Add(-duration)).Error
}

//@author: wuhao
//@function: ClearTable
//@description: 创建数据库表数据
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

func UpdateTable(db *gorm.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	data := make([]system.SysGithub, 1)

	page := "1"
	per_page := "5"
	resp, err := http.Get("https://api.github.com/repos/JiangHaoCode/server-fiber/commits?page=" + page + "&per_page=" + per_page)
	defer func() {
		_ = resp.Body.Close()

	}()
	if err != nil {
		global.LOG.Error("请求Commit错误", zap.Error(err))
	}

	body, _ := io.ReadAll(resp.Body)
	// respData := new([]GithubCommit)
	var respData []system.GithubCommit
	json.Unmarshal(body, &respData)
	time.LoadLocation("Asia/Shanghai")

	for _, val := range respData {
		var temp system.SysGithub
		temp.Author = val.Commit.Author.Name
		temp.CommitTime = val.Commit.Author.Date.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
		temp.Message = val.Commit.Message
		data = append(data, temp)
	}

	return db.Model(&system.SysGithub{}).Create(&data).Error
}

package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	global "server-fiber/model"
	"server-fiber/model/system"
	"sort"
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
	return db.Exec(fmt.Sprintf("update %s SET deleted_at = ? WHERE %s < ?", tableName, compareField), time.Now().Add(-duration), time.Now().Add(-duration)).Error
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

	page := "1"
	per_page := "20"
	resp, err := http.Get("https://api.github.com/repos/JiangHaoCode/server-fiber/commits?page=" + page + "&per_page=" + per_page)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		global.LOG.Error("请求Commit错误", zap.Error(err))
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.LOG.Error("转化失败：", zap.Error(err))
		return err
	}
	var respData []system.GithubCommit
	json.Unmarshal(body, &respData)
	time.LoadLocation("Asia/Shanghai")
	db = db.Model(&system.SysGithub{})
	data := make([]system.SysGithub, 0)
	for _, val := range respData {
		var temp system.SysGithub
		temp.Author = val.Commit.Author.Name
		temp.CommitTime = val.Commit.Author.Date.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
		temp.Message = val.Commit.Message
		data = append(data, temp)
	}
	var dataGithub []system.SysGithub
	for _, item := range data {
		if item.CommitTime != "" {
			db = db.Or("commit_time = ?", item.CommitTime)
		}
	}
	db.Limit(20).Order("id desc").Find(&dataGithub)
	dataInsert := []system.SysGithub{}
	isExist := true
	for _, item := range data {
		for _, itemGithub := range dataGithub {
			if itemGithub.CommitTime == item.CommitTime {
				isExist = false
				break
			} else {
				isExist = true
			}
		}
		if isExist {
			dataInsert = append(dataInsert, item)
			isExist = true
		}
	}
	insertNumber := 0
	if len(dataInsert) != 0 {
		sort.Slice(dataInsert, func(i, j int) bool {
			return timeStrTime(dataInsert[i].CommitTime) < timeStrTime(dataInsert[j].CommitTime)
		})
		for _, item := range dataInsert {
			dataItem := item
			if dataItem.CommitTime != "" {
				insertNumber++
				err = db.Create(&dataItem).Error
			}
		}
	}
	global.LOG.Info(fmt.Sprintf("github 插入 %d 条", insertNumber))

	return err
}

func timeStrTime(valueStr string) int64 {
	loc := time.Local
	fmtStr := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(fmtStr, valueStr, loc)
	return t.Unix()
}

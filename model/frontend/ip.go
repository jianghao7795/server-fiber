package frontend

import "server-fiber/global"

type Ip struct {
	global.MODEL
	Ip        string `json:"ip" form:"ip" gorm:"column:ip;comment:ip;size:50;"`
	ArticleID uint   `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id"`
}

func (Ip) TableName() string {
	return "ips"
}

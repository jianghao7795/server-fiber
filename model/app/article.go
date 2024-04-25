package app

import (
	"server-fiber/global"
	"server-fiber/model/system"
)

type Article struct {
	global.MODEL
	Title           string         `json:"title" form:"title" gorm:"column:title;comment:文章标题;size:191;"`
	Desc            string         `json:"desc" form:"desc" gorm:"column:desc;comment:文章简述;"`
	Content         string         `json:"content" form:"content" gorm:"column:content;comment:文章内容;"`
	State           int            `json:"state" form:"state" gorm:"column:state;comment:文章状态;"`
	UserId          int            `json:"user_id" form:"user_id" gorm:"column:user_id;comment:article 作者id;"`
	Tags            []Tag          `json:"tags" form:"tags" gorm:"many2many:article_tag"` // 多对多 article_tag
	User            system.SysUser `json:"user" form:"user" gorm:"foreignKey:user_id"`
	IsImportant     int            `json:"is_important" form:"is_important" gorm:"column:is_important;comment:首页是否显示;"`
	ReadingQuantity int            `json:"reading_quantity" form:"reading_quantity" gorm:"column:reading_quantity;comment:阅读量;"`
	// Username string `json:"username" form:"username" gorm:"column:username;comment:文章内容;"`
}

// 表名
func (Article) TableName() string {
	return "articles"
}

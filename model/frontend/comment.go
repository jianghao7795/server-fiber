package frontend

import "server-fiber/global"

// Comment 结构体
// 如果含有time.Time 请自行import time包
type Comment struct {
	global.MODEL
	ArticleId int       `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id;size:10;"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleId"`
	ParentId  int       `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:上级;size:10;"`
	Content   string    `json:"content" form:"content" gorm:"column:content;comment:内容;"`
	UserId    int       `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id;"`
	User      User      `json:"user" form:"user" gorm:"foreignKey:UserId"`
	Children  []Comment `json:"children" form:"children" gorm:"foreignKey:ParentId;"`
}

// TableName Comment 表名
func (Comment) TableName() string {
	return "comments"
}

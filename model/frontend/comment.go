package frontend

import "server-fiber/global"

// Comment 结构体
// 如果含有time.Time 请自行import time包
type Comment struct {
	global.MODEL
	ArticleId int     `json:"articleId" form:"articleId" gorm:"column:article_id;comment:文章id;size:10;"`
	Article   Article `json:"article" gorm:"foreignKey:ArticleId"`
	ParentId  int     `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:上级;size:10;"`
	Content   string  `json:"content" form:"content" gorm:"column:content;comment:内容;"`
	UserId    int     `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id;"`
	// UserName  string  `json:"user_name" form:"user_name" gorm:"column:user_name;comment:用户名;"`
	User User `json:"user" form:"user" gorm:"foreignKey:UserId"`
	// UserPraise []system.SysUser `json:"praise" from:"praise" gorm:"many2many:praise"`
	Children []Comment `json:"children" form:"children" gorm:"foreignKey:ParentId;"`
}

// TableName Comment 表名
func (Comment) TableName() string {
	return "comment"
}

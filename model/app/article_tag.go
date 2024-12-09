package app

type ArticleTag struct {
	ArticleId int64 `json:"article_id" form:"article_id" query:"article_id" gorm:"column:article_id;comment:文章id;size:10;"`
	TabId     int64 `json:"tab_id" form:"tab_id" query:"" gorm:"column:app_tab_id;comment:标签id;size:10;"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}

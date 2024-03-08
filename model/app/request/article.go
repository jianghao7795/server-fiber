package request

import (
	"server-fiber/model/app"
	"server-fiber/model/common/request"
)

type ArticleSearch struct {
	app.Article
	TagId string `json:"tag_id" form:"tag_id"`
	request.PageInfo
}

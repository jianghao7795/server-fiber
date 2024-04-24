package request

import (
	"server-fiber/model/app"
	"server-fiber/model/common/request"
)

type ArticleSearch struct {
	app.Article
	request.PageInfo
}

package app

import (
	appService "server-fiber/service/app"
)

// ArticleApi article
type ArticleApi struct{}

var articleService = new(appService.ArticleService)

// BaseMessageApi base message
type BaseMessageApi struct{}

var baseMessageService = new(appService.BaseMessageService)

// CommentApi comment
type CommentApi struct{}

var commentService = new(appService.CommentService)

// TagApi tag
type TagApi struct{}

var appTabService = new(appService.TagService)

// TaskNameApi task
type TaskNameApi struct{}

// FileUploadAndDownloadApi file upload
type FileUploadAndDownloadApi struct{}

var fileUploadService = new(appService.FileUploadService)

// UserApi user
type UserApi struct{}

var userService = new(appService.UserService)

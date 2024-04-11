package app

import (
	appService "server-fiber/service/app"
)

// article
type ArticleApi struct{}

var articleService = new(appService.ArticleService)

// base message
type BaseMessageApi struct{}

var baseMessageService = new(appService.BaseMessageService)

// comment
type CommentApi struct{}

var commentService = new(appService.CommentService)

// tag
type TagApi struct{}

var appTabService = new(appService.TagService)

// task
type TaskNameApi struct{}

// fileupload
type FileUploadAndDownloadApi struct{}

var fileUploadService = new(appService.FileUploadService)

// user
type UserApi struct{}

var userService = new(appService.UserService)

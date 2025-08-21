package main

import (
	_ "server-fiber/docs" // 引入生成的文档
	"server-fiber/core"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	core.RunServer()
}

package main

import (
	"server-fiber/core"
	// jsoniter "github.com/json-iterator/go"
)

/* var json = jsoniter.ConfigCompatibleWithStandardLibrary */

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	core.RunServer()
}

//go:build !windows
// +build !windows

package core

// func initServer(address string, router *fiber.App) server {
// 	s := endless.NewServer(address, router)
// 	s.ReadHeaderTimeout = 20 * time.Second
// 	s.WriteTimeout = 20 * time.Second
// 	s.MaxHeaderBytes = 1 << 20
// 	return s
// }

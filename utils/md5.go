package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// @function: MD5V
// @description: md5加密
// @param: str []byte  char b
// @return: string
func MD5V(str []byte, b ...byte) string {
	h := md5.New() //
	h.Write(str)
	return hex.EncodeToString(h.Sum(b)) // EncodeToString byte切片转换成字符串的编码工具库
}

// @function: MD5VString
func MD5VString(val string, b string) string {
	h := md5.New()
	v := []byte(val)
	d := []byte(b)
	h.Write(v)
	// f4b68e0c8a85ddac35085eb95feb398361fe5c0421922c52dc7797c699664ee13aa4297dc7f20a9cd6615bf000dde6e91cc164988f7c55fc3b4c4c516b8d78c3 /
	return hex.EncodeToString(h.Sum(d)) // EncodeToString byte切片转换成字符串的编码工具库
}

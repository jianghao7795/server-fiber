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

	return hex.EncodeToString(h.Sum(d)) // EncodeToString byte切片转换成字符串的编码工具库
}

package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// sha 加密

func Sha256V(pwd string) string {
	w := sha256.New()
	io.WriteString(w, pwd)
	bw := w.Sum(nil)
	return hex.EncodeToString(bw)
}

func Sha512V(pwd string) string {
	w := sha512.New()

	io.WriteString(w, pwd)
	bw := w.Sum(nil)
	return hex.EncodeToString(bw)
}

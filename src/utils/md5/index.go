package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 字符串 转 MD5
func Md5(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

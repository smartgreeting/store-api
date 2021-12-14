package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// s:加密字符串  sum:密钥
func EncodeMd5(s string, sum []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(sum))
}

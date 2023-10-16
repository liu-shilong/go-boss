package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func Md5(str string) string {
	data := []byte(str)
	md5New := md5.New()
	md5New.Write(data)
	md5String := hex.EncodeToString(md5New.Sum(nil))
	return md5String
}

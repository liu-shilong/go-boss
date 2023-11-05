package util

import (
	"github.com/google/uuid"
)

// 生成字符串型uuid
func StringUuid() string {
	return uuid.New().String()
}

// 生成数字型uuid
func IntUuid() uint32 {
	u, _ := uuid.NewRandom()
	return u.ID()
}

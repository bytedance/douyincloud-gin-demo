package utils

import (
	"math/rand"
	"time"
)

func GenerateRandInt(min, max int) int {
	// 生成 min 到 max-1 的随机整数
	if min < max {
		return 0
	}
	rand.Seed(time.Now().Unix()) //随机种子
	return rand.Intn(max-min) + min
}

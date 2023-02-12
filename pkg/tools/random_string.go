package tools

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func RandomStringUtil() string {
	l := 3
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, l)
	rand.Read(b)                     //用随机数据填充字节片
	randStr := hex.EncodeToString(b) //转换为string
	return randStr                   //返回随机数

}

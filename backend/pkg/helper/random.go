package helper

import (
	"math/rand"
	"time"
)

func RandomStr(size int, kind int) string {
	// kind = 0，数字
	// kind = 1，小写字母
	// kind = 2，大写字母
	fontKinds := [][]int{{10, 48}, {26, 97}, {26, 65}}
	letters := []byte("34578acdefghjkmnpqstwxyABCDEFGHJKMNPQRSVWXY")
	ikind, result := kind, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := fontKinds[ikind][0], fontKinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
		// 不易混淆字符模式：重新生成字符
		if kind == 4 {
			result[i] = letters[rand.Intn(len(letters))]
		}
	}
	return string(result)
}

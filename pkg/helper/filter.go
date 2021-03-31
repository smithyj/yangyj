package helper

import (
	"strconv"
	"strings"
)

func FilterCountryCode(s string) string {
	// 过滤前置无效字符
	for {
		c := s[:1]
		n, _ := strconv.Atoi(c)
		if n > 1 {
			break
		}
		s = strings.TrimLeft(s, "+")
		s = strings.TrimLeft(s, "*")
		s = strings.TrimLeft(s, "0")
	}
	return s
}

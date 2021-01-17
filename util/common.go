package util

import (
	"math/rand"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// HiddenCharacters 隐藏部分字符(中间二分之一）
func HiddenCharacters(s string) string {
	if len(s) <= 2 {
		return s
	}
	b := []byte(s)
	var l = len(b) / 4
	for i := l; i < len(b)-l; i++ {
		b[i] = '*'
	}
	return string(b)
}

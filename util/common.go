/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:common.go
 * Date:2021/4/6 上午11:11
 * Author:sirodeneko
 */

package util

import (
	"math/rand"
	"time"
	"unicode/utf8"
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
	t := utf8.RuneCountInString(s)
	temp := []rune(s)
	if t <= 1 {
		return s
	}
	if t == 2 {
		temp[1] = '*'
		return string(temp)
	}
	if t == 3 {
		temp[1] = '*'
		temp[2] = '*'
		return string(temp)
	}
	var l = len(temp) / 4
	for i := l; i < len(temp)-l; i++ {
		temp[i] = '*'
	}
	return string(temp)
}

func IntToSex(i uint) string {
	if i == 0 {
		return "男"
	} else {
		return "女"
	}
}

// Int64TimeToStr
func Int64TimeToStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006年01月02日")
}

// Int64TimeToStr2
func Int64TimeToStr2(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

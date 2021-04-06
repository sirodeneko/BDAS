/*
 * Copyright (c) 2021.
 * 项目名称:BDAS
 * 文件名称:main.go
 * Date:2021/4/6 上午11:12
 * Author:sirodeneko
 */

package scheduler

import (
	"github.com/panjf2000/ants/v2"
)

const ConcurrentSize = 1000

var pf *ants.PoolWithFunc

func scheduler() {
	pf, _ = ants.NewPoolWithFunc(ConcurrentSize, caFileAndChain, ants.WithPanicHandler(caFileAndChainPanic))
}

func Submit(args interface{}) (err error) {
	if pf == nil {
		scheduler()
	}
	err = pf.Invoke(args)
	return err
}

func Running(args interface{}) (num int) {
	if pf == nil {
		scheduler()
	}
	num = pf.Running()
	return num
}

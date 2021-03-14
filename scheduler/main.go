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

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func changeDuration(d *int32) {
	atomic.AddInt32(d, 5)
	swapped := atomic.CompareAndSwapInt32(d, *d, 30)
	if swapped {
		fmt.Println("new val is ", *d)
	}
}

func main() {
	var cnt int32 = 5
	go changeDuration(&cnt)
	tck := time.NewTicker(time.Duration(time.Millisecond * time.Duration(cnt)))
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-tck.C:
				fmt.Println("executed after ", atomic.LoadInt32(&cnt))
				go changeDuration(&cnt)
				time.Sleep(time.Millisecond * 1)
				tck = time.NewTicker(time.Duration(time.Millisecond *
					time.Duration(atomic.LoadInt32(&cnt))))
			case <-quit:
				tck.Stop()
				fmt.Println("stopping ticker")
			}
		}
	}()

	time.Sleep(time.Second * 1)
	quit <- true
	fmt.Println("updated value ", cnt)
}

//超时机制，select只要其中一个case完成，程序就继续往下执行，不考虑其他case的情况
package main

import (
	"time"
)

timeout := make(chan bool,1)

go func() {
	time.Sleep(1e9)
	timeout <- true
}
select {
	case <-ch:
	case <-timeout:
}

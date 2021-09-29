//利用channel传递特性实现类*nix的管道技术

package main

type PipeData struct {
	value int 
	handler func(int) int
	next chan int
}
func handler(queue chan *PipeData)  {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

go func() {
	time.Sleep(1e9)
	timeout <- true
}
select {
	case <-ch:
	case <-timeout:
}

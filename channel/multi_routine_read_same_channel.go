package main

import (
	"fmt"
	"sync"
)
//多个routine读取同一管道，管道中的每个元素只能为
//一个routine读取，即其为消耗资源。而管道关闭信号
//能够被每一个管道所获取
func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()
		for word := range ch {
			fmt.Println("worker1 : ", word)
		}
	}(&wg, c)

	wg.Add(1)	
	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()
                for word := range ch {
                        fmt.Println("worker2 : ", word)
                }
        }(&wg, c)
	
	c <- 1
	c <- 2
	
	close(c)

	wg.Wait()

}

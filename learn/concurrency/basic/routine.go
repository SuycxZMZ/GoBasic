package basic

import (
	"fmt"
	"time"
)

// 并发计数器结构体
type Counter struct {
	value int
	ch chan int // 定义一个channel
}

// 创建一个并发计数器
func NewCounter() *Counter {
	return &Counter{
		ch : make(chan int, 1), // 定义一个容量为1的channel
	}
}

// 增加计数器的值
func (c *Counter) Increment() {
	c.ch <- 1 // 获取令牌
	c.value++
	<-c.ch // 释放令牌
}

// 获取计数器的值
func (c *Counter) Value() int {
	return c.value
}

// 演示Goroutine泄露问题
func DemoGoroutineLeak() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	// 故意不接收通道数据
	fmt.Println("Goroutine will leak!")
}
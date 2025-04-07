package test

import (
	"learn/concurrency/sync"
	"testing"
	"time"
)

func TestSafeMapConcurrent(t *testing.T) {
	sm := sync.NewSafeMap()
	const workers = 100

	// 并发写入测试
	t.Run("并发写入", func(t * testing.T) {
		// var wg sync.WaitGroup
		// wg.Add(workers)

		for i := 0; i < workers; i++ {
			go func(n int) {
				key := "key"
				sm.Set(key, n)
			}(i)
		}

		// 等待所有goroutine完成写入
		time.Sleep(time.Second)

		val, _ := sm.Get("key")
		if _, ok := val.(int); !ok {
			t.Errorf("类型断言失败")
		}
	})

	t.Run("读写竞争", func(t *testing.T) {
		done := make(chan bool) 
		go func() {
			sm.Set("race","initial")
			done <- true
		}()
		go func ()  {
			sm.Get("race")
			done <- true
		}()

		<-done
		<-done
	})
}
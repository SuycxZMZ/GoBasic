package concurrency

import "time"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// 打包一个result发送到resultChannel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// 从resultChannel接收结果
		result := <-resultChannel
		results[result.string] = result.bool
	}

	time.Sleep(2 * time.Second)
	return results
}

package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://fucker.cn"
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://fucker.cn",
		"http://google.com",
		"http://github.com",
	}

	atcuslResult := CheckWebsites(mockWebsiteChecker, websites)
	want := len(websites)
	got := len(atcuslResult)

	if got != want {
		t.Errorf("want len %v, got len %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://fucker.cn":  false,
		"http://google.com": true,
		"http://github.com": true,
	}

	if !reflect.DeepEqual(atcuslResult, expectedResults) {
		t.Errorf("want %v, got %v", expectedResults, atcuslResult)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

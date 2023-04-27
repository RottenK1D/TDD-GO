package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mocWebsiteChecker(url string) bool {
	if url == "waat://random.fun" {
		return false
	}
	return true
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"waat://random.fun",
	}

	want := map[string]bool{
		"http://google.com":   true,
		"http://facebook.com": true,
		"waat://random.fun":   false,
	}

	got := CheckWebsites(mocWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

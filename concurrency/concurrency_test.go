package concurrency

import (
	"maps"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://badurl.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{"http://google.com", "http://blog.com", "waat://badurl.com"}
	got := CheckWebsites(mockWebsiteChecker, websites)
	assertEqualMaps(t, got, map[string]bool{
		"http://google.com": true,
		"http://blog.com":   true,
		"waat://badurl.com": false,
	})
}

func assertEqualMaps[K comparable, V comparable](t testing.TB, got, want map[K]V) {
	t.Helper()
	if !maps.Equal(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

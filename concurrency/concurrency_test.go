package concurrency

import (
	"maps"
	"testing"
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

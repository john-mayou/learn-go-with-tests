package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url   string
	valid bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))
	resultChannel := make(chan result, len(urls))

	for _, url := range urls {
		go func() {
			resultChannel <- result{url: url, valid: wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.valid
	}

	return results
}

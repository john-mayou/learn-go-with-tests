package blogposts_test

import (
	"errors"
	"io/fs"
	"sort"
	"testing"
	"testing/fstest"

	blogposts "learn-go-with-tests/reading-files/blogposts"

	"github.com/stretchr/testify/assert"
)

type StubFailingFS struct{}

func (s *StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("error")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("returns expected posts", func(t *testing.T) {
		body1 := `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello 1!`
		body2 := `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Hello 2!`

		fs := fstest.MapFS{
			"hello-world.md":   {Data: []byte(body1)},
			"hello-world-2.md": {Data: []byte(body2)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		sort.Slice(posts, func(i, j int) bool { return posts[i].Title < posts[j].Title })

		assert.NoError(t, err)
		assert.Equal(t, []blogposts.Post{
			{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "Hello 1!"},
			{Title: "Post 2", Description: "Description 2", Tags: []string{"rust", "borrow-checker"}, Body: "Hello 2!"},
		}, posts)
	})
	t.Run("returns error on read error", func(t *testing.T) {
		fs := &StubFailingFS{}

		_, err := blogposts.NewPostsFromFS(fs)

		assert.Error(t, err)
	})
}

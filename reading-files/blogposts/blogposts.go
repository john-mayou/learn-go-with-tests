package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSys, ".")
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, file := range dir {
		postFile, err := fileSys.Open(file.Name())
		if err != nil {
			return []Post{}, err
		}
		defer postFile.Close()

		posts = append(posts, buildPost(postFile))
	}

	return posts, nil
}

func buildPost(r io.Reader) Post {
	scanner := bufio.NewScanner(r)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	post := Post{
		Title:       readMetaLine("Title: "),
		Description: readMetaLine("Description: "),
		Tags:        strings.Split(readMetaLine("Tags: "), ", "),
	}

	scanner.Scan() // ignore body separator line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	post.Body = strings.TrimSuffix(buf.String(), "\n")

	return post
}

package blog

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed "posts"
var posts embed.FS

type App struct {
	posts            embed.FS
	slugsToFilenames map[string]string
}

func NewHandler() (http.Handler, error) {
	mux := http.NewServeMux()

	slugsToFilenames, err := parseSlugs(posts)
	if err != nil {
		return nil, fmt.Errorf("failed to parse slugs: %w", err)
	}

	a := App{
		posts:            posts,
		slugsToFilenames: slugsToFilenames,
	}

	mux.HandleFunc("GET /{slug}", a.HandlePost)

	return mux, nil
}

func (a *App) HandlePost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	slugToTitle(slug)

	posts.Open()

}

func (a *App) slugToFile(slug string) (fs.File, error) {
	fileName, ok := a.slugsToFilenames[slug]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	return a.posts.Open(fileName)
}

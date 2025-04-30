package ssr

import (
	"fmt"
	"net/http"
	"path"

	_ "github.com/peterszarvas94/lt2/check"
)

func RunServer() error {
	staticDirPath := path.Join("theme", "static")

	fs := http.FileServer(http.Dir(staticDirPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	faviconPath := path.Join(staticDirPath, "favicon.ico")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, faviconPath)
	})

	indexHandler := &indexHandler{}
	http.Handle("/{$}", indexHandler)

	categoryHandler := &categoryHandler{}
	http.Handle("/category/{category}/{$}", categoryHandler)

	tagHandler := &tagHandler{}
	http.Handle("/tag/{tag}/{$}", tagHandler)

	contentHandler := &contentHandler{}
	http.Handle("/{segments...}", contentHandler)

	fmt.Println("Server is starting on http://localhost:1111")

	err := http.ListenAndServe("localhost:1111", nil)
	if err != nil {
		return err
	}

	return nil
}

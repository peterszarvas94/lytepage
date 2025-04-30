package ssr

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/lt2/fileutils"
	"github.com/peterszarvas94/lt2/pages"
)

type tagHandler struct{}

func (h *tagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tag := r.PathValue("tag")

	tags := fileutils.GetTags()
	files := tags[tag]

	err, pages := pages.GetPages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler := templ.Handler(pages.NotFound())

	if len(files) > 0 {
		handler = templ.Handler(pages.Tag(tag, files))
	}

	handler.ServeHTTP(w, r)
	return
}

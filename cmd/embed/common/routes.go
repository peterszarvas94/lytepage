package common

import (
	"lytepage-test/theme/templates"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/lytepage/fileutils"
)

var CustomRoutes = map[string]templ.Component{
	"/":       templates.CustomIndexPage(fileutils.GetFileByTitle("index")),
	"/search": templates.SearchPage(fileutils.GetFiles()),
	"/docs":   templates.PostsPage(fileutils.GetFiles()),
}

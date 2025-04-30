package common

import (
	"lt2-test/theme/templates"

	"github.com/a-h/templ"
	"github.com/peterszarvas94/lt2/fileutils"
)

var CustomRoutes = map[string]templ.Component{
	"/":       templates.CustomIndexPage(fileutils.GetFileByTitle("index")),
	"/search": templates.SearchPage(fileutils.GetFiles()),
	"/docs":   templates.PostsPage(fileutils.GetFiles()),
}

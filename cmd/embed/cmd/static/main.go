package main

import (
	"fmt"
	"lt2-test/common"
	"os"

	"github.com/peterszarvas94/lt2/custom"
	"github.com/peterszarvas94/lt2/generate"
	"github.com/peterszarvas94/lt2/pages"
	"github.com/peterszarvas94/lt2/static"
)

func main() {
	pages.RegisterPages(common.Pages)

	custom.RegisterCustomRoutes(common.CustomRoutes)

	err := generate.Generate()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = static.RunServer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

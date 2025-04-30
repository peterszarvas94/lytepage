package main

import (
	"fmt"
	"lytepage-test/common"
	"os"

	"github.com/peterszarvas94/lytepage/custom"
	"github.com/peterszarvas94/lytepage/generate"
	"github.com/peterszarvas94/lytepage/pages"
	"github.com/peterszarvas94/lytepage/static"
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

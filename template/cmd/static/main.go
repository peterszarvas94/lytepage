package main

import (
	"fmt"
	"os"
	"scaffhold/common"

	"github.com/peterszarvas94/lytepage/pkg/custom"
	"github.com/peterszarvas94/lytepage/pkg/generate"
	"github.com/peterszarvas94/lytepage/pkg/pages"
	"github.com/peterszarvas94/lytepage/pkg/static"
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

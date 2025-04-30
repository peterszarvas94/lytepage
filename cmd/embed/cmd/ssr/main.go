package main

import (
	"fmt"
	"lt2-test/common"
	"os"

	"github.com/peterszarvas94/lt2/custom"
	"github.com/peterszarvas94/lt2/pages"
	"github.com/peterszarvas94/lt2/ssr"
)

func main() {
	pages.RegisterPages(common.Pages)

	custom.RegisterCustomRoutes(common.CustomRoutes)

	err := ssr.RunServer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

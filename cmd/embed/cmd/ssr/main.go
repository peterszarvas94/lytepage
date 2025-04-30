package main

import (
	"fmt"
	"os"
	"scaffhold/common"

	"github.com/peterszarvas94/lytepage/custom"
	"github.com/peterszarvas94/lytepage/pages"
	"github.com/peterszarvas94/lytepage/ssr"
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

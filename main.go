package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	infoCmd := flag.NewFlagSet("info", flag.ExitOnError)
	infoTemplate := infoCmd.String("template", "", "template")

	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)
	checkoutTemplate := checkoutCmd.String("template", "", "template")
	checkoutDestination := checkoutCmd.String("destination", "", "destination")

	if len(os.Args) < 2 {
		fmt.Println("expected 'list' or 'info' or 'checkout' or 'exit' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "list":
		ReadJSONList()

	case "info":
		infoCmd.Parse(os.Args[2:])
		GetHpTemplateInfo(*infoTemplate)

	case "checkout":
		checkoutCmd.Parse(os.Args[2:6])
		//params := GetHpTemplateParamInfo(*checkoutTemplate)
		var paramMap = make(map[string]string)
		for i, value := range os.Args[6:] {
			if i == 0 || i%2 == 0 {
				res1 := strings.Replace(value, "--", "", 1)
				checkoutparam := checkoutCmd.String(res1, "", res1)
				checkoutCmd.Parse(os.Args[6+i : 6+2+i])
				paramMap[value] = *checkoutparam
			}

		}
		GetTemplateDownload(*checkoutTemplate, *checkoutDestination, paramMap)
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("expected 'list' or 'info' or 'checkout' or 'exit' subcommands")
		os.Exit(1)
	}

}

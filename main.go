package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	infoCmd := flag.NewFlagSet("info", flag.ExitOnError)
	infoTemplate := infoCmd.String("template", "", "template")

	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)
	checkoutTemplate := checkoutCmd.String("template", "", "template")
	checkoutDestination := checkoutCmd.String("destination", "", "destination")
	//checkoutparam1 := checkoutCmd.String("param1", "", "param1")
	//checkoutparam2 := checkoutCmd.String("param2", "", "param2")
	checkoutparam := checkoutCmd.String("param", "", "param")

	//buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	//buildDestination := buildCmd.String("destination", "", "destination")

	if len(os.Args) < 2 {
		fmt.Println("expected 'list' or 'info' or 'checkout' or 'build' or 'exit' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "list":
		readJSONList()

	case "info":
		infoCmd.Parse(os.Args[2:])
		getHpTemplateInfo(*infoTemplate)
	case "checkout":
		checkoutCmd.Parse(os.Args[2:])
		getTemplateDownload(*checkoutTemplate, *checkoutDestination, *checkoutparam)
	// case "build":
	// 	buildCmd.Parse(os.Args[2:])
	// 	exeCommnad(*buildDestination)
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("expected 'list' or 'info' or 'checkout' or 'build' or 'exit' subcommands")
		os.Exit(1)
	}

}

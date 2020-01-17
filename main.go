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


	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	buildDestination := buildCmd.String("destination", "", "destination")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "list":
		readJSONList()

	case "info":
		infoCmd.Parse(os.Args[2:])
		//fmt.Println("subcommand 'info'")
		//fmt.Println("template:", *infoTemplate)
		//fmt.Println(" tail:", infoCmd.Args())
		getHpTemplateInfo(*infoTemplate)
	case "checkout":
		checkoutCmd.Parse(os.Args[2:])
		//fmt.Println("subcommand 'checkout'")
		//fmt.Println("template:", *checkoutTemplate)
		//fmt.Println("destination:", *checkoutDestination)
		//fmt.Println(" tail:", checkoutCmd.Args())
		getTemplateDownload(*checkoutTemplate,*checkoutDestination)
	case "build":
		buildCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'build'")
		fmt.Println("destination:", *buildDestination)
		fmt.Println(" tail:", buildCmd.Args())
		runTemplate(*buildDestination)
    case "exit":
		 os.Exit(1)
	default:
		fmt.Println("expected 'list' or 'info' or 'checkout' or 'build' or 'exit' subcommands")
		os.Exit(1)
	}

}

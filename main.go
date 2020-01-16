package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Print("open")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	suffix := "PS " + dir + ">"
	scanner := bufio.NewScanner(os.Stdin)
	//createJsonFile()
	fmt.Print(suffix + " hp-archetypes --")

	for scanner.Scan() {
		line := scanner.Text()
		s1 := strings.Split(line, "--")
		s2 := s1[0:1]
		//fmt.Print(s1[0:1])
		err := os.Chdir(dir)
		if err != nil {
			fmt.Println("File Path Could not be changed")
		}
		switch strings.Join(s2, "") {
		case "list":
			readJSONList()
			fmt.Print(suffix + " hp-archetypes --")
		case "info":
			getHpTemplateInfo(strings.Join(s1[1:], ""))
			fmt.Print(suffix + " hp-archetypes --")
		case "checkout":
			getTemplateDownload(strings.Join(s1[1:2], ""), strings.Join(s1[2:3], ""))
			fmt.Print(suffix + " hp-archetypes --")
		case "build":
			runTemplate(line, strings.Join(s1[1:], ""))
			//fmt.Print(suffix + " hp-archetypes --")
		case "exit":
			os.Exit(1)
		default:
			fmt.Print(suffix + " hp-archetypes --")

		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

func stringToSlice(str string, splitter string) []string {
	return strings.Split(str, splitter)

}

func sliceToString(slice []string, connector string) string {
	return strings.Join(slice, connector)

}

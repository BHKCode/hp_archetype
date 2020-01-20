package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	fileName    string
	fullUrlFile string
	byteValue   []byte
	archetypes  []Archetype
)

type Archetype struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Param struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Condition string `json:"condition"`
		File      string `json:"file"`
	}
}

func createJSONFile() {
	data := Archetype{
		ID:   1,
		Text: "helper file",
		Type: "hp",
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func readJSONList() {
	openJSONFile("test.json")
	for _, value := range archetypes {
		fmt.Println(value.Type)
	}

}

func getHpTemplateInfo(template string) {
        gopath := os.Getenv("GOPATH")
        abspath := path.Join(gopath, "src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)
	//openJSONFile("test.json")
	//fmt.Println(archetypes)
	//temp := strings.ReplaceAll(template, "template<", "")
	for _, value := range archetypes {
		if value.Type == template {
			fmt.Println("Text : ", value.Text)
			fmt.Println("Type :", value.Type)
			fmt.Println("param name:", value.Param.Name, " type: ", value.Param.Type, " condition:", value.Param.Condition, " file:", value.Param.File)

		}

	}

}

func getTemplateDownload(template string, destination string) {
        gopath := os.Getenv("GOPATH")
        abspath := path.Join(gopath, "src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)
	//openJSONFile("test.json")
	//fmt.Println(archetypes)
	//emp := strings.ReplaceAll(template, "template<", "")
	for _, value := range archetypes {
		if value.Type == template {
			fmt.Println("Text : ", value.Text)
			fmt.Println("Type :", value.Type)
			fmt.Println("param name:", value.Param.Name, "- param Type :", value.Param.Type, "- param condition :",
				value.Param.Condition, "- param file :", value.Param.File)
			fullUrlFile = value.Param.Name
		}

	}
	// Build fileName from fullPath
	buildFileName()
	// Create blank file
	//replacer := strings.NewReplacer("destination<", "", ">", "")
	//output := replacer.Replace(destination)
	file := createFile(destination)
	// Put content on file
	putFile(file, httpClient())

}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullUrlFile)
	checkError(err)
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	defer file.Close()
	checkError(err)
	fmt.Println("Just Downloaded a file %s with size %d", fileName, size)
}

func buildFileName() {
	fileUrl, err := url.Parse(fullUrlFile)
	checkError(err)
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return &client
}

func createFile(path string) *os.File {
	file, err := os.Create(path + string(filepath.Separator) + fileName)
	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
		//fmt.Println("file opening error")
	}
}

func openJSONFile(file string) {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &archetypes)

}

func runTemplate(destination string) {
	//replacer := strings.NewReplacer("destination<", "", ">", "")

	//output := replacer.Replace(destination)

	s1 := strings.Split(destination, string(filepath.Separator))
	sz := len(s1)

	if sz > 0 {
		s1 = s1[:sz-1]
	}
	err := os.Chdir(strings.Join(s1, string(filepath.Separator)))
	if err != nil {
		fmt.Println("File Path Could not be changed")
	}
	exeCommnad(destination)

}

func exeCommnad(destination string) {
	output, err := exec.Command(destination).Output()
	if err == nil {
		fmt.Printf("%s", output)
	} else {
		fmt.Printf("%s", err)
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	fileName    string
	fullUrlFile string
	byteValue   []byte
	archetypes  []Archetype
)

type Archetype struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Param struct {
		Name        string `json:"name"`
		RepoPath    string `json:"repo_path"`
		IncludeGrpc string `json:"include_grpc"`
	}
}

func createJSONFile() {
	data := Archetype{
		ID:   1,
		Name: "helper file",
		Url:  "hp",
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func readJSONList() {
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)
	//openJSONFile("test.json")
	for _, value := range archetypes {
		fmt.Println(value.Name)
	}

}

func getHpTemplateInfo(template string) {
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)

	for _, value := range archetypes {
		if value.Name == template {
			fmt.Println("Name : ", value.Name)
			fmt.Println("Url :", value.Url)
			fmt.Println("param name:", value.Param.Name, " RepoPath: ", value.Param.RepoPath, " IncludeGrpc:", value.Param.IncludeGrpc)

		}

	}

}

func getTemplateDownload(template string, dpath string) {
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)
	//openJSONFile("test.json")
	//fmt.Println(archetypes)
	//emp := strings.ReplaceAll(template, "template<", "")
	for _, value := range archetypes {
		if value.Name == template {
			fmt.Println("Name : ", value.Name)
			fmt.Println("Url :", value.Url)
			fmt.Println("param name:", value.Param.Name, "- RepoPath :", value.Param.RepoPath, "- param condition :",
				value.Param.IncludeGrpc)
			fullUrlFile = value.Url
			fileName = value.Name
		}

	}

	putFile()
	exeCommnad(dpath, fileName)

}

//putFile(file *os.File, client *http.Client)
func putFile() {

	gitPath, err := exec.LookPath("git")
	cmd := exec.Command(fmt.Sprintf("%s", gitPath), "clone", fullUrlFile, fileName)
	err = cmd.Run()
	checkError(err)
	//os.Chdir(fileName)
	//go-archetype transform --transformations=transformations.yml \--source=. \--destination=.tmp/go/my-go-project
	//exeCommnad(dpath)

}

func checkError(err error) {
	if err != nil {
		panic(err)
		//fmt.Println("file opening error")
		log.Fatal(err)
	}
}

func openJSONFile(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &archetypes)

}

func exeCommnad(destination string, filename string) {
	home, _ := os.Getwd()
	goArcPath, err1 := exec.LookPath("go-archetype")
	filepath.Join(destination, filename)
	checkError(err1)
	//cmd := exec.Command("C:\\Go\\bin\\go-archetype.exe", "transform", "--transformations=transformations.yml", "--source=.", "--destination=C:\\Users\\KohaleBh\\Pictures\\gotest", "--", "--ProjectName=abc", "--ProjectDescription=description", "--IncludeReadme=no")
	cmd := exec.Command(fmt.Sprintf("%s", goArcPath), "transform", "--transformations=transformations.yml", "--source=.", "--destination="+destination+string(os.PathSeparator)+filename, "--", "--ProjectName=abc", "--ProjectDescription=description", "--IncludeReadme=yes")
	cmd.Dir = filepath.Join(home, filename)

	err2 := cmd.Run()
	checkError(err2)
}

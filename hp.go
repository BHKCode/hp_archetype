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
		Label1 string `json:"label1"`
		Value1 string `json:"value1"`
		Label2 string `json:"label2"`
		Value2 string `json:"value2"`
	}
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
			//fmt.Println(value)
			fmt.Println("Name : ", value.Name)
			fmt.Println("Url :", value.Url)
			fmt.Println("Params :")
			fmt.Println(value.Param.Label1 + ":" + value.Param.Value1)
			fmt.Println(value.Param.Label2 + ":" + value.Param.Value2)

		}

	}

}

func getTemplateDownload(template string, dpath string, value1 string, value2 string) {
	fmt.Println("tep", dpath)
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)
	for _, value := range archetypes {
		if value.Name == template {

			fullUrlFile = value.Url
			fileName = value.Name
			putFile()
			home, _ := os.Getwd()
			goArcPath, err1 := exec.LookPath("go-archetype")
			if err1 != nil {
				getGoArchetype()
				goArcPath, err1 = exec.LookPath("go-archetype")
			}
			filepath.Join(dpath, fileName)
			checkError(err1)
			//cmd := exec.Command("C:\\Go\\bin\\go-archetype.exe", "transform", "--transformations=transformations.yml", "--source=.", "--destination=C:\\Users\\KohaleBh\\Pictures\\gotest", "--", "--ProjectName=abc", "--ProjectDescription=description", "--IncludeReadme=no")
			cmd := exec.Command(fmt.Sprintf("%s", goArcPath), "transform", "--transformations=transformations.yml", "--source=.", "--destination="+dpath+string(os.PathSeparator)+fileName, "--", "--"+value.Param.Label1+"="+value1, "--"+value.Param.Label2+"="+value2)
			fmt.Println(cmd)
			cmd.Dir = filepath.Join(home, fileName)

			err2 := cmd.Run()
			checkError(err2)
		}

	}

}

//putFile(file *os.File, client *http.Client)
func putFile() {

	gitPath, err := exec.LookPath("git")
	cmd := exec.Command(fmt.Sprintf("%s", gitPath), "clone", fullUrlFile, fileName)
	log.Println(cmd)

	err = cmd.Run()
	checkError(err)

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

func getGoArchetype() {
	goPath, err := exec.LookPath("go")
	cmd := exec.Command(fmt.Sprintf("%s", goPath), "get", "-u", "github.com/rantav/go-archetype")
	log.Println(cmd)
	//log.Println(gitPath)
	err = cmd.Run()
	checkError(err)

}

func getHpTemplateParamInfo(template string) (lab1 string, lab2 string) {
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)

	for _, value := range archetypes {
		if value.Name == template {
			return value.Param.Label1, value.Param.Label2
		}

	}
	return
}

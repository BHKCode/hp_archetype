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
	fullURLFile string
	byteValue   []byte
	archetypes  []Archetype
	Params      []Param
)

const (
	jsonFileName = "test.json"
)

type Archetype struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	URL    string  `json:"url"`
	Params []Param `json:"param"`
}

type Param struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func ReadJSONList() {
	OpenJSONFile()
	//openJSONFile("test.json")
	for _, value := range archetypes {
		fmt.Println(value.Name)
	}

}

func GetHpTemplateInfo(template string) {
	OpenJSONFile()
	for _, value := range archetypes {
		if value.Name == template {
			//fmt.Println(value)
			fmt.Println("Name : ", value.Name)
			fmt.Println("Url :", value.URL)
			fmt.Println("Params :")
			for _, pValue := range value.Params {
				fmt.Println(pValue.Label + ":" + pValue.Value)
			}
		}

	}

}

func GetTemplateDownload(template string, dpath string, param map[string]string) {
	OpenJSONFile()
	for _, value := range archetypes {
		if value.Name == template {

			fullURLFile = value.URL
			fileName = value.Name
			PutFile()
			home, _ := os.Getwd()
			goArcPath, err1 := exec.LookPath("go-archetype")
			if err1 != nil {
				GetGoArchetype()
				goArcPath, err1 = exec.LookPath("go-archetype")
			}
			filepath.Join(filepath.FromSlash(dpath), fileName)

			CheckError(err1)
			args := make([]string, 10)
			i := 1
			args[0] = "transform"
			args[1] = "--transformations=transformations.yml"
			args[2] = "--source=."
			args[3] = "--destination=" + dpath + string(os.PathSeparator) + fileName
			args[4] = "--"
			for key, value := range param {
				args[4+i] = key + "=" + value
				i = i + 1
			}
			//cmd := exec.Command(fmt.Sprintf("%s", goArcPath), "transform", "--transformations=transformations.yml", "--source=.", "--destination=C:\\Users\\KohaleBh\\Pictures\\gotest", "--", "--ProjectName=abc", "--ProjectDescription=description", "--IncludeReadme=no")
			//cmd := exec.Command(fmt.Sprintf("%s", goArcPath), "transform", "--transformations=transformations.yml", "--source=.", "--destination="+dpath+string(os.PathSeparator)+fileName, "--", paramStr)
			cmd := exec.Command(fmt.Sprintf("%s", goArcPath), args...)
			fmt.Println(cmd)
			cmd.Dir = filepath.Join(home, fileName)

			err2 := cmd.Run()
			CheckError(err2)
		}

	}

}
func PutFile() {
	gitPath, err := exec.LookPath("git")
	cmd := exec.Command(fmt.Sprintf("%s", gitPath), "clone", fullURLFile, fileName)
	log.Println(cmd)

	err = cmd.Run()
	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func OpenJSONFile() {
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+jsonFileName)
	jsonFile, err := os.Open(abspath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &archetypes)

}

func GetGoArchetype() {
	goPath, err := exec.LookPath("go")
	cmd := exec.Command(fmt.Sprintf("%s", goPath), "get", "-u", "github.com/rantav/go-archetype")
	log.Println(cmd)
	//log.Println(gitPath)
	err = cmd.Run()
	CheckError(err)

}

func GetHpTemplateParamInfo(template string) []string {
	OpenJSONFile()
	lables := make([]string, 10)
	for _, value := range archetypes {
		if value.Name == template {
			for i, values := range value.Params {
				lables[i] = values.Label
			}
			return lables
		}
	}
	return nil
}

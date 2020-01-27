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
		Label2 string `json:"label2"`
		Label3 string `json:"label3"`
		Value1 string `json:"value1"`
		Value2 string `json:"value2"`
		Value3 string `json:"value3"`
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
			fmt.Println("param Label1:", value.Param.Label1, " Label2: ", value.Param.Label2, " Label3:", value.Param.Label3)
			fmt.Println("param Value1:", value.Param.Value1, " Value2: ", value.Param.Value2, " Value3:", value.Param.Value3)

		}

	}

}

func getTemplateDownload(template string, dpath string, value1 string, value2 string, value3 string) {
	//fmt.Println(archetypes)
	gopath := os.Getenv("GOPATH")
	abspath := path.Join(gopath, string(os.PathSeparator)+"src"+string(os.PathSeparator)+"github.com"+string(os.PathSeparator)+"BHKCode"+string(os.PathSeparator)+"hp_archetype"+string(os.PathSeparator)+"test.json")
	openJSONFile(abspath)
	//openJSONFile("test.json")
	//fmt.Println(archetypes)
	//emp := strings.ReplaceAll(template, "template<", "")
	for _, value := range archetypes {
		if value.Name == template {
			//fmt.Println("Name : ", value.Name)
			//fmt.Println("Url :", value.Url)
			//fmt.Println("param Label1:", value.Param.Label1, " Label2: ", value.Param.Label2, " Label3:", value.Param.Label3)
			//fmt.Println("param Value1:", value.Param.Value1, " Value2: ", value.Param.Value2, " Value3:", value.Param.Value3)
			fullUrlFile = value.Url
			fileName = value.Name
			putFile()
			home, _ := os.Getwd()
			goArcPath, err1 := exec.LookPath("go-archetype")
			filepath.Join(dpath, fileName)
			checkError(err1)
			//cmd := exec.Command("C:\\Go\\bin\\go-archetype.exe", "transform", "--transformations=transformations.yml", "--source=.", "--destination=C:\\Users\\KohaleBh\\Pictures\\gotest", "--", "--ProjectName=abc", "--ProjectDescription=description", "--IncludeReadme=no")
			cmd := exec.Command(fmt.Sprintf("%s", goArcPath), "transform", "--transformations=transformations.yml", "--source=.", "--destination="+dpath+string(os.PathSeparator)+fileName, "--", "--"+value.Param.Label1+"="+value1, "--"+value.Param.Label2+"="+value2, "--"+value.Param.Label3+"="+value3)
			//fmt.Println(cmd)
			cmd.Dir = filepath.Join(home, fileName)

			err2 := cmd.Run()
			checkError(err2)
		}

	}

}

//putFile(file *os.File, client *http.Client)
func putFile() {

	//gitPath, err := exec.LookPath("git")
	//cmd := exec.Command(fmt.Sprintf("%s", gitPath), "clone", fullUrlFile, fileName)
	cmd := exec.Command("git", "clone", fullUrlFile, fileName)
	log.Println(cmd)
	//log.Println(gitPath)
	err := cmd.Run()
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

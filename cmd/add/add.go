package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type Information struct {
	id      string
	name    string
	title   string
	dirname string
	path    string
}

var rootPath string

func mapStr(ss []string, fn func(string) string) []string {
	mapped := make([]string, len(ss))
	for i, s := range ss {
		mapped[i] = fn(s)
	}
	return mapped
}

func generatePath(i *Information) {
	splitted := mapStr(strings.Split(strings.TrimSpace(i.name), " "), strings.ToLower)
	i.title = strings.Join(mapStr(splitted, strings.Title), " ")
	i.dirname = i.id + "-" + strings.Join(splitted, "-")
	i.path = path.Join("problems", i.dirname)

}

func writeFile(i Information, tpl *template.Template, name string) {

	var data bytes.Buffer
	err := tpl.ExecuteTemplate(&data, name+".tmpl", nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		ioutil.WriteFile(path.Join(i.path, name)+".go", data.Bytes(), 0777)
	}

}

func generateFiles(i Information) {
	tpl, err := template.ParseGlob("cmd/add/templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	writeFile(i, tpl, "main")
	writeFile(i, tpl, "case")
	writeFile(i, tpl, "solution")
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	if !strings.HasSuffix(path, "leetcode-go") {
		fmt.Println("This command should executed from project root directory")
		os.Exit(0)
	}
	rootPath = path
}

func main() {
	info := Information{}

	getId(&info)
	getName(&info)
	generatePath(&info)

	filePaths := map[string]string{
		"main":     path.Join(info.path, "main.go"),
		"solution": path.Join(info.path, "solution.go"),
		"case":     path.Join(info.path, "case.go"),
	}

	fmt.Printf(
		"+ leetcode-go/%s\n+ leetcode-go/%s\n+ leetcode-go/%s\n",
		filePaths["main"],
		filePaths["solution"],
		filePaths["case"],
	)

	confirm()
	os.Mkdir(info.path, 0755)
	generateFiles(info)
}

package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/maxia51/golog"
)

func main() {

	path := flag.String("dir", ".", "location you want to scann")
	flag.Parse()

	files, err := ioutil.ReadDir(*path)

	if err != nil {
		golog.E(err.Error())
		return
	}

	files = cleanFiles(&files)

	createJSON(&files)

}

func movieToJSON(s string) string {
	//func Split(s, sep string) []string

	str := strings.Split(s, ".")
	str = append(str[:len(str)-1], str[len(str):]...)

	s = strings.Join(str, ".")

	return s
}

func cleanFiles(files *[]os.FileInfo) []os.FileInfo {

	var tmp []os.FileInfo

	for i := 0; i < len(*files); i++ {

		extension := filepath.Ext((*files)[i].Name())
		if extension == ".mkv" || extension == ".mp4" {
			golog.T((*files)[i].Name())
			tmp = append(tmp, (*files)[i])
		}
	}

	return tmp
}

func createJSON(files *[]os.FileInfo) {
	json, err := os.Create("movies.json")

	if err != nil {
		golog.F(err.Error())
		return
	}

	json.WriteString("[")
	for i := 0; i < len(*files); i++ {

		json.WriteString("{ \"name\":\"" + movieToJSON((*files)[i].Name()) + "\" }")

		if i != len(*files)-1 {
			json.WriteString(",")
		}
	}
	json.WriteString("]")
}

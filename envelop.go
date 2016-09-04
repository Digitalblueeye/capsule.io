package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

var sourceFile string
var value map[string]string

func init() {
	value = make(map[string]string)

	fileinfo, _ := ioutil.ReadDir(basepath)
	for _, file := range fileinfo {
		if !file.IsDir() && strings.Contains(file.Name(), ".envelop") {
			File, _ := ioutil.ReadFile(basepath + "/" + file.Name())
			sourceFile = string(File)
		}
	}
}

func main() {
	fmt.Println(Env("apple"))
}

func Env(key string) interface{} {
	if len(value) == 0 {
		load()
	}
	return value
}

func SetSource(filepath string) {
	//sourceFile = os.NewFile(nil, filepath)
	File, _ := ioutil.ReadFile(filepath)
	sourceFile = string(File)
}

/* TODO func DisableMemoryCache(do bool) {

}
*/

func load() {
	inKey := true
	currentKey := ""
	currentVal := ""
	index := 0
	for i, char := range sourceFile {
		fmt.Println(string(char), i)
		if inKey && string(char) == "=" || inKey && string(char) == ":" {
			inKey = false
			currentKey = sourceFile[index:i]
			index = i + 1
			fmt.Println(currentKey)
		} else if !inKey {
			currentVal = sourceFile[index:i]
			fmt.Println(currentVal, "ik")
			inKey = true
			value["{`key`"+currentKey+"}"] = "{`val`" + currentVal + "}"
		}

	}
}

//Parse file line by line,

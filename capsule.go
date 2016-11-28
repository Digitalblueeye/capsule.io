package capsuleio

import (

	// "fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	storage map[string]string
	source  string
)

//Sets the source file for the key value store (aka the capsule)
func Open(capsule string) {
	File, _ := ioutil.ReadFile(capsule)
	source = string(File)
	load()
}

//Gets the string value from the current capsule
func Get(key string) string {
	if len(storage) < 1 {
		//load local file from directory
		storage = make(map[string]string)

		basepath, _ := os.Getwd()

		fileinfo, _ := ioutil.ReadDir(basepath)
		for _, file := range fileinfo {
			if !file.IsDir() && strings.Contains(file.Name(), ".capsule") {
				//fmt.Println(basepath + "/" + file.Name())
				File, _ := ioutil.ReadFile(basepath + "/" + file.Name())
				source = string(File)
				break
			}
		}
		load()
		return storage[key]
	} else {
		return storage[key]
	}
	return "Missing a capsule file? Incorrect formatting?"
}

func load() {
	lines := strings.Split(source, "\n")
	for i := range lines {
		group := strings.SplitAfterN(lines[i], "=", 2)
		if len(group) >= 2 {
			storage[strings.TrimSpace(strings.Replace(group[0], "=", "", 1))] = strings.TrimSpace(group[1])
		}
	}
}

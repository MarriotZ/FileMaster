package _case

import (
	"log"
	"os"
	"path"
)

func ReadWriteFiles() {
	list := getFiles(sourceDir)
	for _, f := range list {
		bytes, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		_, name := path.Split(f)
		destName := destDir + "normal/" + name
		err = os.WriteFile(destName, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

package _case

import (
	"io"
	"log"
	"os"
	"path"
)

func CopyDirToDir() {
	list := getFiles(sourceDir)
	for _, f := range list {
		_, name := path.Split(f)
		destFileName := destDir + "copy/" + name
		CopyFile(f, destFileName)
	}
}

func CopyFile(srcName, destName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

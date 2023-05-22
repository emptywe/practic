package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if err := filepath.Walk("./sample", func(path string, info os.FileInfo, errr error) error {
		if info.IsDir() || filepath.Ext(info.Name()) == ".DS_Store" {
			return nil
		}
		rgxp := regexp.MustCompile("[A-Za-z0-9()]+")
		fmt.Println(path)
		fmt.Println(path[0 : len(path)-len(info.Name())])
		newName := strings.Join(rgxp.FindAllString(info.Name()[0:len(info.Name())-len(filepath.Ext(info.Name()))], -1), "")

		//fmt.Println(strings.Replace(path, "+", "_", -1))
		if err := os.Rename(path, path[0:len(path)-len(info.Name())]+newName+filepath.Ext(info.Name())); err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}
}

package tools

import (
	"fmt"
	"os"
	"strings"
)

//获取当前路径

func GetPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Path := strings.Split(dir, "/cmd")[0]
	return Path

}

func PublishVideoToPublic(video []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("create file fail")
		return err
	}
	defer file.Close()

	_, err = file.Write(video)
	if err != nil {
		fmt.Println("write file fail")
		return err
	}

	return nil

}

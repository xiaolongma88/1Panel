package test

import (
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/app/service"
	"log"
	"os"
	"testing"
)

func TestGetContent(t *testing.T) {
	config := service.CameraService{}
	content, err := config.GetContent()
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	fmt.Printf("%+v\n", content)
}

func TestGetImages(t *testing.T) {
	dir, err := os.Open("/jetsonDetect/res/")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]string, len(files))
	for _, file := range files {
		result = append(result, file.Name())
	}
	fmt.Println(result)
}

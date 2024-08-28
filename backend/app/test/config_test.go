package test

import (
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/app/service"
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

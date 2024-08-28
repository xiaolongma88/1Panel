package service

import (
	"encoding/json"
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/app/dto/response"
	"github.com/1Panel-dev/1Panel/backend/utils/files"
)

type CameraService struct {
}

type ICameraService interface {
	GetContent() (response.Config, error)
}

func NewICameraService() ICameraService {
	return &CameraService{}
}

func (f *CameraService) GetContent() (response.Config, error) {
	// 读取 JSON 文件
	fileOp := files.NewFileOp()
	var data, err = fileOp.GetContent("/jetsonDetect/config/config.json")
	if err != nil {
		fmt.Println("读取文件错误：", err)
		return response.Config{}, err
	}

	var config = response.Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("解析配置文件错误：", err)
		return response.Config{}, err
	}
	fmt.Println("解析成功")
	// 复写配置
	return config, nil
}

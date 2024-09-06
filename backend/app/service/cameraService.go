package service

import (
	"encoding/json"
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/app/dto/response"
	"github.com/1Panel-dev/1Panel/backend/utils/files"
	"io/fs"
	"os"
	"os/exec"
)

type CameraService struct {
}

type ICameraService interface {
	GetContent() (response.Config, error)
	UpdateContent(config dto.CameraContent) error
	GetImages() ([]string, error)
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

func (f *CameraService) UpdateContent(config dto.CameraContent) error {
	newContent, err := json.Marshal(config)
	if err != nil {
		fmt.Println("转换出错:", err)
		return err
	}
	fileOp := files.NewFileOp()
	if err := fileOp.SaveFile("/jetsonDetect/config/config.json", string(newContent), fs.FileMode(0755)); err != nil {
		fmt.Println("保存出错", err)
		return err
	}
	// 定义shell脚本的路径
	scriptPath := "/jetsonDetect/jetson_restart.sh"
	// 创建一个Cmd对象，执行shell脚本
	cmd := exec.Command("/bin/bash", scriptPath)
	// 获取命令的输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// 打印输出结果
	fmt.Println(string(output))
	return nil
}

func (f *CameraService) GetImages() ([]string, error) {
	dir, err := os.Open("/jetsonDetect/res/")
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(files))
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result, nil
}

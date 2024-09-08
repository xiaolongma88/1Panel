package service

import (
	"encoding/json"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/app/dto/response"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/1Panel-dev/1Panel/backend/utils/files"
	"io/fs"
	"os"
	"os/exec"
	"path"
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
	fileOp := files.NewFileOp()
	var data, err = fileOp.GetContent(global.CONF.DirConfig.AppConfig)
	if err != nil {
		return response.Config{}, err
	}

	var config = response.Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return response.Config{}, err
	}
	return config, nil
}

func (f *CameraService) UpdateContent(config dto.CameraContent) error {
	newContent, err := json.Marshal(config)
	if err != nil {
		return err
	}
	fileOp := files.NewFileOp()
	if err := fileOp.SaveFile(global.CONF.DirConfig.AppConfig, string(newContent), fs.FileMode(0755)); err != nil {
		return err
	}
	scriptPath := path.Join(global.CONF.DirConfig.AppDir, "jetson_restart.sh")
	cmd := exec.Command("/bin/bash", scriptPath)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	global.LOG.Infof(string(output))
	return nil
}

func (f *CameraService) GetImages() ([]string, error) {
	dir, err := os.Open(global.CONF.DirConfig.ResultDir)
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

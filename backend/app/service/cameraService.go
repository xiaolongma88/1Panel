package service

import (
	"encoding/json"
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/app/dto/response"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/1Panel-dev/1Panel/backend/utils/files"
	"github.com/go-acme/lego/v4/log"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
)

type CameraService struct {
}

type ICameraService interface {
	GetContent() (response.Config, error)
	UpdateContent(config dto.CameraContent) error
	GetImages(camId string) ([]string, error)
	ParseRTSP(rtsp dto.RtspInfo) error
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

func (f *CameraService) GetImages(camId string) ([]string, error) {
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
		imageName := strings.Split(file.Name(), "-")
		if imageName[0] == camId {
			result = append(result, file.Name())
		}
	}
	return result, nil
}

func (f *CameraService) ParseRTSP(rtspInfo dto.RtspInfo) error {
	cmdOne := exec.Command("pkill", "ffmpeg")
	if err := cmdOne.Run(); err != nil {
		fmt.Println("Error closing ffmpeg processes:", err)
	}
	fmt.Println("All ffmpeg processes closed successfully.")
	cmdStr := fmt.Sprintf("ffmpeg -i %s -c copy -f flv rtmp://localhost:8888/live/%s", rtspInfo.RtspAddr, rtspInfo.CameraID)
	cmd := exec.Command("sh", "-c", cmdStr)
	if err := cmd.Start(); err != nil {
		log.Printf("Error starting FFmpeg command: %v", err)
		return err
	}
	fmt.Printf("Started pushing RTSP stream with ID %s to Nginx", rtspInfo.CameraID)
	return nil
}

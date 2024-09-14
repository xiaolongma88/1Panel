package v1

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/constant"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags Camera
// @Summary get camera config info
// @Description 获取相机配置
// @Accept json
// @Success 200
// @Security ApiKeyAuth
// @Router /camera/conf/info [get]
func (b *BaseApi) GetCameraConfig(c *gin.Context) {
	config, err := cameraService.GetContent()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, config)
}

// @Tags Camera
// @Summary edit camera config info
// @Description 全量更新相机配置
// @Accept json
// @Success 200
// @Security ApiKeyAuth
// @Router /camera/conf/update [post]
func (b *BaseApi) UpdateCameraConfig(c *gin.Context) {
	var req dto.CameraContent
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := cameraService.UpdateContent(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

func (b *BaseApi) GetImages(c *gin.Context) {
	result, err := cameraService.GetImages()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, result)
}

func (b *BaseApi) ShowImages(c *gin.Context) {
	imageName := c.Query("imageName")
	c.File(global.CONF.DirConfig.ResultDir + imageName)
}

func (b *BaseApi) ShowVideos(c *gin.Context) {
	videoName := c.Query("videoName")
	c.File(global.CONF.DirConfig.TestVideo + videoName)
}

func (b *BaseApi) ShowRTSP(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	cameraService.ParseRTSP(c, id)
}

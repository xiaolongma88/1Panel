package v1

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/constant"
	"github.com/gin-gonic/gin"
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

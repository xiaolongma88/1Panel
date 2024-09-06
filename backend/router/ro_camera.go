package router

import (
	v1 "github.com/1Panel-dev/1Panel/backend/app/api/v1"
	"github.com/1Panel-dev/1Panel/backend/middleware"
	"github.com/gin-gonic/gin"
)

type CameraRouter struct{}

func (c *CameraRouter) InitRouter(Router *gin.RouterGroup) {
	cameraRouter := Router.Group("camera")
	cameraRouter.Use(middleware.JwtAuth()).Use(middleware.SessionAuth()).Use(middleware.PasswordExpired())
	baseApi := v1.ApiGroupApp.BaseApi
	{
		cameraRouter.GET("/conf/info", baseApi.GetCameraConfig)
		cameraRouter.POST("/conf/update", baseApi.UpdateCameraConfig)
		cameraRouter.GET("/conf/img", baseApi.GetImages)
		cameraRouter.GET("/config/res", baseApi.ShowImages)
	}
}

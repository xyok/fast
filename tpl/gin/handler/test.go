package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary 服务连通性测试
// @Description 服务连通性测试接口
// @Tags 测试
// @Success 200 {object} schema.Response
// @Failure 400 {object} schema.Response
// @BasePath /
// @Router /ping [get]
func PingApi(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

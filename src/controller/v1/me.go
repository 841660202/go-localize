package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @Summary 测试接口
// @Description 描述信息
// @Router /me [get]
func GetMe(c *gin.Context) {
	c.JSON(http.StatusOK, Res{
		Code: 200,
		Data: nil,
		Msg:  "GetMe: hello world, I am gin",
	})
}
func PostMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "PostMe: hello world, I am gin",
	})
}
func PutMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "PutMe: hello world, I am gin",
	})
}
func DeleteMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "DeleteMe: hello world, I am gin",
	})
}

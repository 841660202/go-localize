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
// @ID Helloworld
// @Description do ping
// @Tags 例子
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetHelloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @Summary 获取用户
// @Description 描述信息
// @Tags 用户
// @ID GetMe
// @Router /me [get]
func GetMe(c *gin.Context) {
	c.JSON(http.StatusOK, Res{
		Code: 200,
		Data: nil,
		Msg:  "GetMe: hello world, I am gin",
	})
}

// @Summary 新增用户
// @Description 描述信息
// @Tags 用户
// @ID PostMe
// @Router /me [post]
func PostMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "PostMe: hello world, I am gin",
	})
}

// @Summary 更新用户
// @Description 描述信息
// @Tags 用户
// @ID PutMe
// @Router /me [put]
func PutMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "PutMe: hello world, I am gin",
	})
}

// @Summary 删除用户
// @Description 描述信息
// @Tags 用户
// @ID DeleteMe
// @Router /me [delete]
func DeleteMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "DeleteMe: hello world, I am gin",
	})
}

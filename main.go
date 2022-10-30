package main

import (
	"fmt"
	"go-localize/src/router"
	"os"

	gin_swagger_knife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8001
// @BasePath /ap1/v1
func main() {
	router := router.StartRouter()
	// 获取 swag int 命令生成的swagger.json文件里的内容
	swaggerJson := getFileContent("./docs/swagger.json")
	//初始化gin路径
	gin_swagger_knife.InitSwaggerKnife(router, swaggerJson)
	router.Run(":8001")
}

func getFileContent(fpath string) string {
	bytes, err := os.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return string(bytes)
}

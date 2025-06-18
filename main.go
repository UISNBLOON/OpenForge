// 主程序入口文件，可在此初始化依赖、设置路由并启动 HTTP 服务器
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func createProjectStructure() {
	dirs := []string{"handlers", "models", "router", "config", "static"}
	files := []string{
		"config/config.toml",
		"handlers/storage_handler.go",
		"models/file_model.go",
		"router/routes.go",
	}

	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
	}

	for _, file := range files {
		os.WriteFile(filepath.Join(file), []byte("package "+filepath.Base(filepath.Dir(file))), 0644)
	}
}

func main() {
	createProjectStructure()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":5244")
}

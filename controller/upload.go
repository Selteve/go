package controller

import (
	"fmt"
	"time"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)


/*
 * @Author: <EMAIL>
 * @Version: 1.0.0
 * @Date: 2018/6/5
 * @Description: 文件上传，分为单个文件和多个文件上传，
 */

// 单个文件上传
func UploadFile(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})	
		return
	}
	// 获取文件名
	filename := file.Filename
	// 获取文件后缀
	ext := strings.Split(filename, ".")[1]
	// 生成时间戳作为文件名
	timestamp := time.Now().Unix()
	// 拼接文件名
	filename = fmt.Sprintf("%d.%s", timestamp, ext)
	// 判断文件类型，上传到指定目录
	if ext == "jpg" || ext == "png" {
		err := c.SaveUploadedFile(file, "./files/images/"+filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"url": "http:127.0.0.1:8000/files/images/" + filename,
		})
	} else if ext == "mp4" || ext == "avi" || ext == "rmvb" || ext == "flv" || ext == "wmv" {
		err := c.SaveUploadedFile(file, "./files/video/"+filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"url": "http:127.0.0.1:8000/files/video/" + filename,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "文件类型错误",
		})
		return
	}
	
}
// 多个文件上传
func UploadFiles() {
}
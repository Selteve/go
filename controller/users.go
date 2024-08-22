package controller

import (
	"fmt"
	"time"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	Config "gitee.com/under-my-umbrella/cloud/config"
	UserDb "gitee.com/under-my-umbrella/cloud/db/users"
	Utils "gitee.com/under-my-umbrella/cloud/utils"
)


// 用户注册
func UsersRegister(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    file, err := c.FormFile("user_img")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
        return
    }
	// 生成时间戳作为新文件名
	timestamp := time.Now().Unix()
	extension := filepath.Ext(file.Filename) // 获取文件扩展名
	newFilename := fmt.Sprintf("%d%s", timestamp, extension)
	
	// 生成哈希密码
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 调用数据库，必须先判断用户是否存在
	if UserDb.CheckUser(username) {
		// 用户已存在
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "用户已存在",
		})
		return
	} else {
		// 保存文件到指定目录
		err = c.SaveUploadedFile(file, "./files/avatar/" + newFilename)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to save file")
			return
		}
		UserImg := "http://127.0.0.1:8000/files/avatar/" + newFilename
		// 继续注册,对密码进行加密
		if UserDb.CreateUser(username, hashedPassword, UserImg) {
			// 注册成功
			c.JSON(200, gin.H{
				"code": 1,
				"msg": "注册成功",
			})
		}
	}
}

// 用户登录
func UserLogin(c *gin.Context) {
	// 绑定json数据
	var user Config.UserConfig
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// 判断用户是否存在
	if UserDb.CheckUser(user.Username) {
		// 判断密码是否正确
		if UserDb.GetUser(user.Username, []byte(user.Password)) {
			// 查找用户信息
			id, err := UserDb.GetUserIdByUsername(user.Username)
			if err != nil {
				return 
			}
			// 生成token
			token, err := Utils.GenerateToken(id, user.Username)
			if err != nil {
				return
			}
			c.JSON(200, gin.H{
				"code": 1,
				"msg": "登录成功",
				"token": token,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "密码错误",
			})
		}
	} else {
		// 用户不存在
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "用户不存在",
		})
	}
}
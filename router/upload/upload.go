package upload
import (
	"net/http"
	"github.com/gin-gonic/gin"
	Middleware "gitee.com/under-my-umbrella/cloud/middleware"
)
// SetupUploadRoutes setup upload routes
func SetupUploadRoutes(api *gin.RouterGroup) {
	upload := api.Group("/upload")
	upload.Use(Middleware.VrifyToken()) // 验证token
	upload.GET("", getAllUpload)
	upload.GET("/:id", getAllUpload)
	upload.POST("", getAllUpload)
	upload.PUT("/:id", getAllUpload)
	upload.DELETE("/:id", getAllUpload)
}

func getAllUpload(c *gin.Context) {
    // 从上下文中获取用户信息
    ID, existsID := c.Get("id")
    username, existsUsername := c.Get("username")

    if !existsID || !existsUsername {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "用户信息未找到",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":  "pong",
        "ID":   ID,
        "username": username,
    })
}
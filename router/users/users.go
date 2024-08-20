package users
import (
	"net/http"
	"github.com/gin-gonic/gin"
	Controller "gitee.com/under-my-umbrella/cloud/controller"
)
// SetupUserRoutes sets up all route handling for the users resource
func SetupUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")

	users.GET("", getAllUpload)
	users.GET("/:id", getAllUpload)
	users.POST("/register", Controller.UsersRegister) // 注册
	users.POST("/login", Controller.UserLogin)        // 登录
	users.PUT("/:id", getAllUpload)
	users.DELETE("/:id", getAllUpload)
}

func getAllUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

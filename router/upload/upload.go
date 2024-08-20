package upload
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// SetupUploadRoutes setup upload routes
func SetupUploadRoutes(api *gin.RouterGroup) {
	upload := api.Group("/upload")

	upload.GET("", getAllUpload)
	upload.GET("/:id", getAllUpload)
	upload.POST("", getAllUpload)
	upload.PUT("/:id", getAllUpload)
	upload.DELETE("/:id", getAllUpload)
}

func getAllUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
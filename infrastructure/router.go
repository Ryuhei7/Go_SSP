package infrastructure

import (
	"github.com/gin-gonic/gin"
	//"ryuhei/Go_ssp/interface/controllers"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	// エンドポイントが出来次第ここでグループにして設定
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	return r

}

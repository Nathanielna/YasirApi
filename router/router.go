package router
import (
    "github.com/gin-gonic/gin"
    ."yasirapi/apis"
)
func InitRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/", IndexApi)
    router.GET("/baidutop", GetBaidutopsApi)
    router.GET("/weibohot", GetWeibohotsApi)
    return router
}
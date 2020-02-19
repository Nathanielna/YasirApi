package router
import (
    "github.com/gin-gonic/gin"
    ."yasirapi/apis"
)
func InitRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/baidutop", GetPersonsApi)
    return router
}
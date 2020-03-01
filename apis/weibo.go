package apis

import (
    "net/http"
    "log"
    // "fmt"
    // "strconv"
    "github.com/gin-gonic/gin"
     ."yasirapi/models"
)


func GetWeibohotsApi(c *gin.Context) {
    var p WeiboHot
    weibohots, err := p.GetWeiboHots()
    if err != nil {
        log.Fatalln(err)
    }

    c.JSON(http.StatusOK, gin.H{
        "data": weibohots,
    })

}


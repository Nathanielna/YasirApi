package apis

import (
    "net/http"
    "log"
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
     ."yasirapi/models"
)

func IndexApi(c *gin.Context) {
    c.String(http.StatusOK, "It works")
}

func AddBaidutopApi(c *gin.Context) {
    title := c.Request.FormValue("Title")
    sort := c.Request.FormValue("Sort")

    p := Baidutop{Title: title, Sort: sort}

    ra, err := p.AddBaidutop()
    if err != nil {
        log.Fatalln(err)
    }
    msg := fmt.Sprintf("insert successful %d", ra)
    c.JSON(http.StatusOK, gin.H{
        "msg": msg,
    })
}

func GetBaidutopsApi(c *gin.Context) {
    var p Baidutop
    baidutops, err := p.GetBaidutops()
    if err != nil {
        log.Fatalln(err)
    }

    c.JSON(http.StatusOK, gin.H{
        "data": baidutops,
    })

}

func GetBaidutopApi(c *gin.Context) {
    cid := c.Param("id")
    id, err := strconv.Atoi(cid)
    if err != nil {
        log.Fatalln(err)
    }
    p := Baidutop{Id: id}
    baidutop, err := p.GetBaidutop()
    if err != nil {
        log.Fatalln(err)
    }

    c.JSON(http.StatusOK, gin.H{
        "baidutop": baidutop,
    })

}

func ModBaidutopApi(c *gin.Context) {
    cid := c.Param("id")
    id, err := strconv.Atoi(cid)
    if err != nil {
        log.Fatalln(err)
    }
    p := Baidutop{Id: id}
    err = c.Bind(&p)
    if err != nil {
        log.Fatalln(err)
    }
    ra, err := p.ModBaidutop()
    if err != nil {
        log.Fatalln(err)
    }
    msg := fmt.Sprintf("Update person %d successful %d", p.Id, ra)
    c.JSON(http.StatusOK, gin.H{
        "msg": msg,
    })
}

func DelBaidutopApi(c *gin.Context) {
    cid := c.Param("id")
    id, err := strconv.Atoi(cid)
    if err != nil {
        log.Fatalln(err)
    }
    p := Baidutop{Id: id}
    ra, err := p.DelBaidutop()
    if err != nil {
        log.Fatalln(err)
    }
    msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
    c.JSON(http.StatusOK, gin.H{
        "msg": msg,
    })
}




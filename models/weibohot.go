package models
import (
   //"log"
   db "yasirapi/databases"
)
//定义person类型结构
type WeiboHot struct {
   Id        int    `json:"id"`
   Title  string    `json:"title"`
   Sort  string `json:"order"`
}

type WeiboHotResult struct {
  
   Title  string    `json:"title"`
   Sort  string `json:"order"`
}


func (p *WeiboHot) GetWeiboHots() (weibohots []WeiboHotResult, err error) {
   weibohots = make([]WeiboHotResult, 0)
   rows, err := db.SqlDB.Query("SELECT wb_title,wb_order FROM weibohot")
   defer rows.Close()
   if err != nil {
      return
   }
   for rows.Next() {
      var weibohot WeiboHotResult
      rows.Scan( &weibohot.Title, &weibohot.Sort)
      weibohots = append(weibohots, weibohot)
   }
   if err = rows.Err(); err != nil {
      return
   }
   return
}


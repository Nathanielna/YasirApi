package models
import (
   "log"
   db "yasirapi/databases"
)
//定义person类型结构
type Baidutop struct {
   Id        int    `json:"id"`
   Title  string    `json:"title"`
   Time  string `json:"time"`
   Sort  string `json:"sort"`
}

func (p *Baidutop) AddBaidutop() (id int64, err error) {
   rs, err := db.SqlDB.Exec("INSERT INTO baidutop(title, sort) VALUES (?, ?)", p.Title, p.Sort)
   if err != nil {
      return
   }
   id, err = rs.LastInsertId()
   return
}

func (p *Baidutop) GetBaidutops() (baidutops []Baidutop, err error) {
   baidutops = make([]Baidutop, 0)
   rows, err := db.SqlDB.Query("SELECT title,sort FROM baidutop")
   defer rows.Close()
   if err != nil {
      return
   }
   for rows.Next() {
      var baidutop Baidutop
      rows.Scan( &baidutop.Title, &baidutop.Sort)
      baidutops = append(baidutops, baidutop)
   }
   if err = rows.Err(); err != nil {
      return
   }
   return
}

func (p *Baidutop) GetBaidutop() (baidutop Baidutop, err error) {
   err = db.SqlDB.QueryRow("SELECT id, title, sort FROM baidutop WHERE id=?", p.Id).Scan(
      &baidutop.Id, &baidutop.Title, &baidutop.Sort,
   )
   return
}

func (p *Baidutop) ModBaidutop() (ra int64, err error) {
   stmt, err := db.SqlDB.Prepare("UPDATE baidutop SET title=?, sort=? WHERE id=?")
   defer stmt.Close()
   if err != nil {
      return
   }
   rs, err := stmt.Exec(p.Title,p.Sort,p.Id,)
   if err != nil {
      return
   }
   ra, err = rs.RowsAffected()
   return
}

func (p *Baidutop) DelBaidutop() (ra int64, err error) {
   rs, err := db.SqlDB.Exec("DELETE FROM baidutop WHERE id=?", p.Id)
   if err != nil {
      log.Fatalln(err)
   }
   ra, err = rs.RowsAffected()
   return
}

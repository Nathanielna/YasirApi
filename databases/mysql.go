package databases
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)
//因为我们需要在其他地方使用SqlDB这个变量，所以需要大写代表public
var SqlDB *sql.DB
//初始化方法
func init() {
    var err error
    SqlDB, err = sql.Open("mysql", "root:yyyxxxppp@tcp(139.9.57.122:3306)/look_around?parseTime=true")
    if err != nil {
        log.Fatal(err.Error())
    }
    //连接检测
    err = SqlDB.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }
}
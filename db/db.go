package db

import(
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	"github.com/yang-f/beauty/settings"
	"github.com/yang-f/beauty/utils/log"
)

func Query(sql string, params ...interface{})([]mysql.Row, mysql.Result, error) {
    // db := mysql.New("tcp", "", settings.MYSQL.Url, settings.MYSQL.Username, settings.MYSQL.Userpass, settings.MYSQL.Database)
    db := mysql.New("tcp", "", settings.Mysql["mysql_host"], settings.Mysql["mysql_user"], settings.Mysql["mysql_pass"], settings.Mysql["mysql_database"])
	if err := db.Connect(); err != nil {
        return nil, nil, err
    }
    defer db.Close()
    log.Printf("query:"+sql,params...)
    raws, res, err := db.Query(sql,params...)
    if err != nil {
    	return nil, nil, err
    }

    return raws, res, err
}
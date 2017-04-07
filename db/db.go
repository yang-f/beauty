package db

import(
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/thrsafe"
    "github.com/yang-f/beauty/settings"
    "github.com/yang-f/beauty/utils/log"
)

func Query(sql string, params ...interface{})([]mysql.Row, mysql.Result, error) {
    db := mysql.New("tcp", "", settings.Local["mysql_host"], settings.Local["mysql_user"], settings.Local["mysql_pass"], settings.Local["mysql_database"])
    if err := db.Connect(); err != nil {
        log.Println(err)
        return nil, nil, err
    }
    defer db.Close()
    log.Printf("query:"+sql,params...)
    raws, res, err := db.Query(sql,params...)
    if err != nil {
        log.Println(err)
        return nil, nil, err
    }

    return raws, res, err
}
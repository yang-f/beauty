package settings

import (
    "log"
    "encoding/json"
    "io/ioutil"
)

var Listen = ":8080"

var HmacSampleSecret = []byte("beauty")

var LogFile = "/var/log/beauty/beauty.log"

var SettingsFile = "/srv/filestore/settings/latest.json"

var Domain = "beauty.com"

var Mysql = map[string]string{}

func init(){
    bytes, err := ioutil.ReadFile(SettingsFile)
    if err != nil {
        log.Println("ReadFile: ", err.Error())
        return
    }
 
    if err := json.Unmarshal(bytes, &Mysql); err != nil {
        log.Println("Unmarshal: ", err.Error())
        return
    }
}
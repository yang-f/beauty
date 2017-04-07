package settings

import (
	"log"
	"encoding/json"
	"io/ioutil"
)

var Listen = ":8080"

var HmacSampleSecret = []byte("hulu")

var LogFile = "/var/log/beauty/beauty.log"

var Domain = "xxxx.com"

var Local = map[string]string{}

func init(){
	bytes, err := ioutil.ReadFile("/srv/filestore/settings/latest.json")
	if err != nil {
		log.Println("ReadFile: ", err.Error())
		return
	}
 
	if err := json.Unmarshal(bytes, &Local); err != nil {
		log.Println("Unmarshal: ", err.Error())
		return
	}
}
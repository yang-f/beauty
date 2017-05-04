package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Listen = ":8080"

var HmacSampleSecret = []byte("whatever")

var LogFile = "/var/log/beauty/beauty.log"

var Domain = "xxxx.com"

var DefaultOrigin = "http://origin.com"

var Local = map[string]string{}

var TrimArgs = "whatever"

func init() {
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

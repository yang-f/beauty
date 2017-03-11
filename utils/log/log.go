package log

import(
	"log"
	"os"
    "github.com/yang-f/beauty/settings"
    "runtime"
    "fmt"
    "strings"
)

var Fatalln = log.Fatalln
var Fatal 	= log.Fatal

func Printf(format string, params ...interface{}){
	_,f,line,_ := runtime.Caller(1)
    log.Printf(format, params ...)
	file, err := os.OpenFile(settings.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Printf("%v",err)
		return
	}
    defer file.Close()
    _, err = file.Seek(0, os.SEEK_END)
    if err != nil {
    	return
    }
    args := strings.Split(f, "/")
    f = args[len(args)-1]
    msg := fmt.Sprintf("%v:%v(%v)", line, format, f)
    logger := log.New(file, "", log.LstdFlags)
    logger.Printf(msg, params ...)
}

func Println(v ...interface{}) {
	_,f,line,_ := runtime.Caller(1)
    log.Println(v ...)
	file, err := os.OpenFile(settings.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Printf("%v",err)
		return
	}
    defer file.Close()
    _, err = file.Seek(0, os.SEEK_END)
    if err != nil {
    	return
    }
    args := strings.Split(f, "/")
    f = args[len(args)-1]
    msg := fmt.Sprintf("%v:%v(%v)", line, fmt.Sprintln(v ...), f)
    logger := log.New(file, "", log.LstdFlags)
    logger.Println(msg)
}




// MIT License

// Copyright (c) 2017 FLYING

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package log

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/yang-f/beauty/settings"
)

var Fatalln = log.Fatalln
var Fatal = log.Fatal

var ticker = time.NewTicker(time.Minute * 60 * 12)

func init() {
	go func() {
		for _ = range ticker.C {
			archive()
		}
	}()
}

func Printf(format string, params ...interface{}) {
	_, f, line, _ := runtime.Caller(1)
	log.Printf(format, params...)
	file, err := os.OpenFile(settings.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Printf("%v", err)
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
	logger.Printf(msg, params...)
}

func Println(v ...interface{}) {
	_, f, line, _ := runtime.Caller(1)
	log.Println(v...)
	file, err := os.OpenFile(settings.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer file.Close()
	_, err = file.Seek(0, os.SEEK_END)
	if err != nil {
		return
	}
	args := strings.Split(f, "/")
	f = args[len(args)-1]
	msg := fmt.Sprintf("%v:%v(%v)", line, fmt.Sprintln(v...), f)
	logger := log.New(file, "", log.LstdFlags)
	logger.Println(msg)
}

func shortFileName(file string) string {
	return filepath.Base(file)
}

func shortFileDir(file string) string {
	return filepath.Dir(file)
}

func archive() error {
	info, _ := os.Stat(settings.LogFile)
	if info.Size() > 1024*1024*50 {
		target := fmt.Sprintf("%v.%v.tar.gz",
			shortFileName(settings.LogFile),
			time.Now().Format("2006-01-02-15-04"),
		)
		tmp := fmt.Sprintf("%v.%v.tmp",
			shortFileName(settings.LogFile),
			time.Now().Format("2006-01-02-15-04"),
		)
		in := bytes.NewBuffer(nil)
		cmd := exec.Command("sh")
		cmd.Stdin = in
		go func() {
			in.WriteString(fmt.Sprintf("cd %v\n", shortFileDir(settings.LogFile)))
			in.WriteString(fmt.Sprintf("cp %v %v\n", shortFileName(settings.LogFile), tmp))
			in.WriteString(fmt.Sprintf("echo '' > %v\n", shortFileName(settings.LogFile)))
			in.WriteString(fmt.Sprintf("tar -czvf %v %v\n", target, tmp))
			in.WriteString(fmt.Sprintf("rm %v\n", tmp))
			in.WriteString("exit\n")
		}()
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

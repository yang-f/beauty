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

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yang-f/beauty/router"
	"github.com/yang-f/beauty/settings"
	"github.com/yang-f/beauty/utils"
	"github.com/yang-f/beauty/utils/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("beauty", "A command-line tools of beauty.")
	demo     = app.Command("demo", "Demo of web server.")
	generate = app.Command("generate", "Generate a new app.")
	name     = generate.Arg("name", "AppName for app.").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case generate.FullCommand():
		GOPATH := os.Getenv("GOPATH")
		appPath := fmt.Sprintf("%s/src/%s", GOPATH, *name)
		origin := fmt.Sprintf("%s/src/github.com/yang-f/beauty/etc/demo.zip", GOPATH)
		dst := fmt.Sprintf("%s.zip", appPath)
		_, err := utils.CopyFile(dst, origin)
		if err != nil {
			fmt.Println(err.Error())
		}
		utils.Unzip(dst, appPath)
		os.RemoveAll(dst)
		helper := utils.ReplaceHelper{
			Root:    appPath,
			OldText: "{appName}",
			NewText: *name,
		}
		helper.DoWrok()
		log.Printf("Generate %s success.", *name)
	case demo.FullCommand():
		log.Printf("Start server on port %s", settings.Listen)
		router := router.New(router.BRoutes)
		log.Fatal(http.ListenAndServe(settings.Listen, router))
	}
}

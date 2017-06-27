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

package utils

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func CurrentPath() string {
	path, _ := exec.LookPath(os.Args[0])
	return path
}

func Rand() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func MD5(msg string, upper bool) string {
	h := md5.New()
	h.Write([]byte(msg))
	cipherStr := h.Sum(nil)
	if upper {
		return strings.ToUpper(hex.EncodeToString(cipherStr))
	}
	return hex.EncodeToString(cipherStr)
}

func Post(url string, params string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "*/*")
	req.Header.Set("connection", "Keep-Alive")
	req.Header.Set("user-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;SV1)")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Unzip(src_zip string, dest string) error {
	unzip_file, err := zip.OpenReader(src_zip)
	if err != nil {
		return err
	}
	os.MkdirAll(dest, 0755)
	for _, f := range unzip_file.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			_, err = io.Copy(f, rc)
			if err != nil {
				if err != io.EOF {
					return err
				}
			}
			f.Close()
		}
	}
	unzip_file.Close()
	return nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

type ReplaceHelper struct {
	Root    string
	OldText string
	NewText string
}

func (h *ReplaceHelper) DoWrok() error {
	return filepath.Walk(h.Root, h.walkCallback)
}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		return nil
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(buf)
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)
	ioutil.WriteFile(path, []byte(newContent), 0)
	return err
}

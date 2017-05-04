package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
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
	} else {
		return hex.EncodeToString(cipherStr)
	}
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

func Trim(r rune) bool {
	return strings.Contains("()nume", string(r))
}

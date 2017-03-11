package utils

import (
    "net/http"
    "encoding/json"
    "github.com/yang-f/beauty/utils/log"
)


type response struct{
    Status int `json:"status"`
    Description string `json:"description"`
    Code string `json:"code"`
}

func Response(w http.ResponseWriter, description string,code string, status int) {
    out := &response{status, description, code}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    log.Printf("response:\t%s",description);
    w.WriteHeader(status)
    w.Write(b)
}

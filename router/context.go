package router

import (
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/yang-f/beauty/models"
	"github.com/yang-f/beauty/utils"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

type handler func(c *Context) error

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		W: w,
		R: r,
	}
	log.Printf("ser beauty")
	if e := fn(c); e != nil {
		log.Printf("info:%+v", e)
		switch e := errors.Cause(e).(type) {
		case *models.APPError:
			utils.Response(c.W, e.Message, e.Code, e.Status)
		default:
			utils.Response(c.W, "未知错误", "UnknownErr", 200)
		}
	}
}

func (inner handler) corsHeader() handler {
	return handler(func(c *Context) error {
		c.W.Header().Set("Access-Control-Allow-Origin", "*")
		c.W.Header().Set("Access-Control-Allow-Credentials", "true")
		c.W.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		c.W.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
		inner.ServeHTTP(c.W, c.R)
		return nil
	})
}

func (inner handler) logger() handler {
	return handler(func(c *Context) error {
		start := time.Now()
		inner.ServeHTTP(c.W, c.R)
		log.Printf(
			"%s\t%s\t%s",
			c.R.Method,
			c.R.RequestURI,
			time.Since(start),
		)
		return nil
	})
}

func CorsHeader(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
		inner.ServeHTTP(w, r)
	})
}

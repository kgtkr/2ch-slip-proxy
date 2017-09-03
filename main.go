package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			res, _ := http.Get("http://ipcheck.mycurse.net/")
			defer res.Body.Close()
			byteArray, _ := ioutil.ReadAll(res.Body)
			ip := string(byteArray)
			now := time.Now()
			hasher := md5.New()
			hasher.Write([]byte(fmt.Sprintln(ip, r.Host, now.Year(), now.Month(), now.Day())))
			hash := hex.EncodeToString(hasher.Sum(nil))
			r.Header.Set("User-Agent", "Monazilla/1.00 ("+hash+"/1)")

			return r, nil
		})

	log.Fatal(http.ListenAndServe(":8080", proxy))
}

package main

import (
	"SSO/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync/atomic"
)

func main() {
	var parsedLbBackendsArr []*url.URL
	var counter atomic.Int64

	cfg, errLoadConf := config.Load()
	if errLoadConf != nil {
		log.Fatal("load config failed : ", errLoadConf)
	}
	lbBackends := cfg.LbBackends
	lbBackendsArr := strings.Split(lbBackends, ",")

	for _, backendServer := range lbBackendsArr {
		backendServerParsed, errParsed := url.Parse(backendServer)
		if errParsed != nil {
			log.Fatal("parse url failed : ", errParsed)
		}
		parsedLbBackendsArr = append(parsedLbBackendsArr, backendServerParsed)

	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			n := counter.Add(1)
			i := (n - 1) % int64(len(parsedLbBackendsArr))
			target := parsedLbBackendsArr[i]
			r.SetURL(target)
			r.SetXForwarded()
		},
	}
	if err := http.ListenAndServe(":8080", proxy); err != nil {
		log.Fatal(err)
	}
}

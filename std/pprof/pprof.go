package main

import (
	"net/http"
	"net/http/pprof"
)

func main() {
	c := make(chan bool)
	go func() {
		profServeMux := http.NewServeMux()
		profServeMux.HandleFunc("/debug/pprof/", pprof.Index)
		profServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		profServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		profServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		err := http.ListenAndServe(":7789", profServeMux)
		if err != nil {
			panic(err)
		}
		c <- true
	}()
	<-c
}

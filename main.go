package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("rec req:%v", req.RequestURI)
		var body string
		if bs, err := ioutil.ReadAll(req.Body); err == nil {
			body = string(bs)
		}

		resp := map[string]interface{}{
			"method":      req.Method,
			"proto":       req.Proto,
			"request_uri": req.RequestURI,
			"remote_addr": req.RemoteAddr,
			"raw_query":   req.URL.RawQuery,
			"host":        req.Host,
			"path":        req.URL.Path,
			"raw_url":     req.URL.RawPath,
			"header":      req.Header,
			"body":        body,
		}
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "err:%v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Write(data)
	})

	log.Println("listen on 8080")
	log.Fatalf("fail to serve, err:%v", http.ListenAndServe(":8080", mux))
}
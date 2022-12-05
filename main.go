package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Handler is the entry point for fission function
func Handler(w http.ResponseWriter, r *http.Request) {
	subpath := r.Header["X-Fission-Params-Subpath"]
	requestURI := "/" + strings.Join(subpath, ",")
	switch requestURI {
	case "/":
		writeData(w, http.StatusOK, "text/plain; charset=utf-8", []byte("pong"))
	case "/ping":
		writeData(w, http.StatusOK, "text/plain; charset=utf-8", []byte("pong"))
	case "/cctv.test":
		var epgList = make([]string, 0)
		for _, c := range cctv {
			epgList = append(epgList, c.epg)
		}
		resp, err := fastTest(epgList)
		if err != nil {
			writeData(w, http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
			return
		}

		writeJson(w, http.StatusOK, resp)
	case "/cctv.m3u":
		var m3uList = make([]string, 0)
		for _, c := range cctv {
			m3uList = append(m3uList, c.m3u)
		}
		resp, err := fastGet(m3uList)
		if err != nil {
			writeData(w, http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
			return
		}
		writeData(w, http.StatusOK, resp.contentType, resp.body)
	case "/cctv.xml":
		var epgList = make([]string, 0)
		for _, c := range cctv {
			epgList = append(epgList, c.epg)
		}
		resp, err := fastGet(epgList)
		if err != nil {
			writeData(w, http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
			return
		}
		writeData(w, http.StatusOK, resp.contentType, resp.body)
	}

}

func writeData(w http.ResponseWriter, code int, contentType string, data []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
	w.Write(data)
}

func writeJson(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. `

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func serveLoremIpsum(w http.ResponseWriter, r *http.Request) {
	// Send the initial headers saying we're gonna stream the response.
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	index := 0

	for i := 0; i < *size; i++ {
		for j := 0; j < 1024; j++ {
			for k := 0; k < 1024; k++ {
				if index == len(loremIpsum) {
					index = 0
				}
				fw.Write([]byte(string(loremIpsum[index])))
				index++
			}
		}
	}
}

var port = flag.Int("port", 9999, "port number")
var size = flag.Int("size", 10, "size in mega bytes")

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveLoremIpsum)

	err := http.ListenAndServe(":"+strconv.Itoa(*port), mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

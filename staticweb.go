package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", staticFileHandler)

	http.ListenAndServe(":8085", nil)
}

// staticファイルを転送するハンドラ
func staticFileHandler(w http.ResponseWriter, r *http.Request) {

	fname := "." + r.RequestURI
	file, ferr := os.OpenFile(fname, os.O_RDONLY, 0600)
	defer file.Close()

	if ferr != nil {
		log.Println("not found file:" + fname)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reader := bufio.NewReader(file)

	buf := make([]byte, 1024, 1024)
LOOP:
	for {
		n, err := reader.Read(buf)
		switch {
		case err == nil:
			w.Write(buf[:n])
		case err == io.EOF:
			log.Println("static file=" + fname)
			break LOOP
		case err != nil:
			log.Printf("file error occured. err=%s\n", err)
			break LOOP
		}
	}

	return
}

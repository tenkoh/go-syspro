package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// gzip writer の出力先を http.ResponseWriterにする
// json encoder の出力先をgzip writerとos.Stdoutにする
// encode したらgzip writerはフラッシュしてやる

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	writer := gzip.NewWriter(w)
	defer writer.Close()
	multi := io.MultiWriter(writer, os.Stdout)
	encoder := json.NewEncoder(multi)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
	writer.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	//some broweser sends this request, and handler is executed twice. this prevents the twice-operation.
	http.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":8080", nil)
}

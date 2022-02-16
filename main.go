package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"main/lib/manip"
	"net/http"
)

func board(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fields := map[string]string{}
	for k, v := range params {
		if k == "title" {
			continue
		}
		fields[k] = v[0]
	}

	buf := new(bytes.Buffer)
	im, err := manip.StringStuff(params.Get("title"), fields)
	if err != nil {
		http.Error(w, "Failed To Generate", http.StatusInternalServerError)
		return
	}
	err = png.Encode(buf, im)
	if err != nil {
		http.Error(w, "Failed To Generate", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}

func main() {
	http.HandleFunc("/board", board)
	fmt.Printf("Go Get Some\nListening on port 8080\n")
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatal(err)
	}
}

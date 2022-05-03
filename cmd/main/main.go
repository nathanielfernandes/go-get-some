package main

import (
	"fmt"
	"log"
	"main/lib/helpers"
	"main/lib/manip"
	"net/http"
)

func discordpass(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fields := map[string]string{}
	for k, v := range params {
		fields[k] = v[0]
	}

	im, err := manip.DiscordPass(fields)

	if err != nil {
		http.Error(w, "Failed To Generate", http.StatusInternalServerError)
		return
	}

	helpers.ImageResponse(&im, w)
}

func board(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fields := map[string]string{}
	for k, v := range params {
		if k == "title" {
			continue
		}
		fields[k] = v[0]
	}

	im, err := manip.StringStuff(params.Get("title"), fields)

	if err != nil {
		http.Error(w, "Failed To Generate", http.StatusInternalServerError)
		return
	}

	helpers.ImageResponse(&im, w)
}

func main() {
	http.HandleFunc("/board", board)
	http.HandleFunc("/passport", discordpass)

	fmt.Printf("Go Get Some\nListening on port 80\n")
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatal(err)
	}
}

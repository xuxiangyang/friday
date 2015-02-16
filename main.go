package main

import (
	"encoding/json"
	"friday"
	"log"
	"net/http"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		date := friday.NewDate(time.Now())
		jsonStr, err := json.Marshal(date)
		if err != nil {
			log.Fatal(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonStr)
	}
}

func main() {
	http.HandleFunc("/", requestHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

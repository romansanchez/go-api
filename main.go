package main

import (
  "net/http"
  "log"
)

func main() {
	router := ApiRouter()
	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
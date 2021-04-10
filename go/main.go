package main

import (
	"net/http"
	"with_b/api"
)

func main() {
	api.Test()
	
	http.HandleFunc("/lecture", api.Lectures)
	http.HandleFunc("/review", api.Review)

	http.ListenAndServe(":3030", nil)
}

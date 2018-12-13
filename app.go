package main

import (
	"net/http"
	"restfulDemo/routes"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":9090", r)
}

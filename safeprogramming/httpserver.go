package main

import (
	"net/http"
)

func main(){
	h:=http.FileServer(http.Dir("."))
	http.ListenAndServe(":8081",h)
}
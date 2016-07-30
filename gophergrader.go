package main

import (
	"fmt"
	"github.com/blockninja/ninjarouter"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func main() {

}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

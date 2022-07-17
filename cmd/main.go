package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Midle(myF))

	http.ListenAndServe(":3000", nil)
}

func myF(w http.ResponseWriter, r *http.Request) {
	log.Println("you are fuck")
	w.WriteHeader(http.StatusOK)
}


func Midle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("Midle")
		next(w, r)
	}
}

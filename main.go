package main

import (
	"ChallengeBox/service"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {

	ConnectDatabase()

	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/challenge", service.Challenges)

	http.ListenAndServe(":3000", nil)

}

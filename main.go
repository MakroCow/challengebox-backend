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

var port = ":3000"

func main() {

	ConnectDatabase()

	fmt.Println("Successfully started up on Port", port)

	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/challenge", service.Challenges)
	http.HandleFunc("/tag/", service.Tag)
	http.HandleFunc("/tag/challengeId/", service.TagById)
	http.HandleFunc("/setup", service.Setup)

	http.ListenAndServe(port, nil)

}

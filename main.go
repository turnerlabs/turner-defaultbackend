package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	var port string

	port = getenv("PORT", "8001")
	fmt.Println("Listening on : " + port)

	r := mux.NewRouter()

	r.HandleFunc(getenv("HEALTHCHECK", "/healthz"), healthCheck)
	r.HandleFunc("/{blah:.*}", defaultRoute)
	r.HandleFunc("/", defaultRoute)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", r)

	if os.Getenv("ENABLE_LOGGING") == "true" {
		http.ListenAndServe(":"+port, loggedRouter)
	} else {
		http.ListenAndServe(":"+port, nil)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	var x = "{\"Hello\":" + "\"" + os.Getenv("PRODUCT") + "-" + os.Getenv("ENVIRONMENT") + "\"}"
	w.Write([]byte(x))
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Default Backend", http.StatusNotImplemented)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

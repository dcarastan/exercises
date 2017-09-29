package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Run(":8888")
}

// http://localhost:8888/?password=secret123

// MyMiddleware func
func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there..." + r.URL.Query().Get("password"))

	password := r.URL.Query().Get("password")
	if password == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized:"+password, 401)
	}

	log.Println("Logging on the way back...")
}

package main

import (
	"fmt"
	"net/http"
)

func middlewareMethodValidation(next http.Handler, allowedMethods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowedMethods {
			if method == r.Method {
				next.ServeHTTP(w, r)
				return
			}

		}
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/trending", middlewareMethodValidation(http.HandlerFunc(handleTrending), "GET"))
	mux.Handle("/post", middlewareMethodValidation(http.HandlerFunc(handleCreatePost), "POST"))

	err := http.ListenAndServe("0.0.0.0:1234", mux)
	if err != nil {
		panic(err.Error())
	}
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Thanks for yout POST")
}

func handleTrending(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Today Trending is : Toilet Binus lt.4 ")
}

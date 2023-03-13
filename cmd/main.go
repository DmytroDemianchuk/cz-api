package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmytrodemianchuk/cz-api/router"
	"github.com/gorilla/handlers"
)

func main() {
	// cors
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST)
	// handler := cors.Default().Handler(mux)
	// http.ListenAndServe(":8080", handler)

	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")

	// listening at port
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"}),
		)(r)))
	fmt.Println("Listening at port 8080 ...")
}

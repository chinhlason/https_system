package main

import (
	"encoding/json"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	http.Handle("/api/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := json.Marshal(map[string]string{"message": "hello world"})
		w.Write(resp)
	}))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
	})

	log.Println("Server started at :2999")
	err := http.ListenAndServeTLS(":2999", "./ssl/public.crt", "./ssl/private.key", c.Handler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

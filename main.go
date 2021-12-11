package main

import (
	"discord-bot/bot"
	"discord-bot/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = bot.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}).Methods("GET")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, r))

}

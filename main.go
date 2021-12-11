package main

import (
	"discord-bot/bot"
	"discord-bot/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	// fmt.Println("port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

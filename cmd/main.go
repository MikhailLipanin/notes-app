package main

import (
	"github.com/MikhailLipanin/notes-app"
	"github.com/MikhailLipanin/notes-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(notes.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

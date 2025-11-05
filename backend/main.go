package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vaidik-bajpai/Nexus/backend/internal/db"
	"github.com/vaidik-bajpai/Nexus/backend/internal/handler"
	"github.com/vaidik-bajpai/Nexus/backend/internal/store"
)

func main() {
	port := flag.Int("port", 8080, "server port address")
	flag.Parse()

	dbClient := db.NewDB()
	defer dbClient.Disconnect()

	store := store.NewStore(dbClient)
	hdl := handler.NewHandler(store)
	mux := hdl.SetupRoutes()

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	log.Printf("server started on port %d\n", *port)
	if err := srv.ListenAndServe(); err != nil {
		log.Println("server failed to start ", err)
		os.Exit(1)
	}
}

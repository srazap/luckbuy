package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/srazap/luckbuy/db"
	"github.com/srazap/luckbuy/router"
)

func main() {

	// initialize viper
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// auto migrate database
	db.MigrateModels()

	// create api server
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server will remain active")

	r := mux.NewRouter()

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      handlers.LoggingHandler(os.Stdout, router.HandleRoutes(r)),
	}

	// handle routes

	log.Println("Starting server on port 8080")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("ListenAndServe Error::", err.Error())
		}
	}()
	log.Println("ListenAndServe() on Port 8080")
	log.Println("Press Ctrl + C to shutdown the server gracefully...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)

	log.Println("Shutting down API server")
	os.Exit(0)

}

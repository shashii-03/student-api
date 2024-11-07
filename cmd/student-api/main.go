package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/te-shashikant/student-api/internal/config"
	"github.com/te-shashikant/student-api/internal/http/handlers/students"
)


func main() {
	//load config
	cfg:= config.MustLoad()
	//database setup
	//setup routing
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students",students.New())
	//start server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	slog.Info("Server started at", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func ()  {
		
		err := server.ListenAndServe() 
		if err !=nil{
			log.Fatal("Server not started")
		}
	}()

	<-done

	//gracefull shutdown
	slog.Info("Shutting down server")
		ctx, cancel:= context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		
	if	err:=server.Shutdown(ctx); err!=nil{
		slog.Error("failed to shutdown", slog.String("error", err.Error()))
	}

		slog.Info("Shutdown success")


} 
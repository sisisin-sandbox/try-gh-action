package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serverMux := http.NewServeMux()
	serverMux = buildServerMux(serverMux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Server is running on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %s\n", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")

	if err := srv.Shutdown((ctx)); err != nil {
		fmt.Printf("Shutdown(): %s\n", err)
	}

	fmt.Println("Server is shutdown")
}

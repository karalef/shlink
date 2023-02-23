package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/karalef/shlink/app"
	"github.com/karalef/shlink/cmd"
	"github.com/karalef/shlink/urlservice"
	"google.golang.org/grpc"
)

func init() {
	log.SetPrefix("shlink: ")
}

func main() {
	log.Println("starting...")

	var application *app.App

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go cmd.ListenSig(ctx, func(s os.Signal) {
		cancel()
		log.Println(s.String(), "shutting down...")
		if application != nil {
			application.Shutdown(10 * time.Second)
		}
		os.Exit(0)
	})

	serviceAddr := cmd.Getenv("SERVICE_ADDR", "api:9898")
	conn, err := grpc.DialContext(ctx, serviceAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial %s: %s", serviceAddr, err)
	}
	defer conn.Close()

	log.Println("established connection to", serviceAddr)

	srv := urlservice.NewService(conn)
	appURL := cmd.Getenv("APP_URL", "http://localhost:8080")
	application, err = app.New(srv, appURL)
	if err != nil {
		log.Fatalln("failed to create app:", err)
	}

	log.Printf("app %s started", appURL)

	err = application.Run(":" + cmd.Getenv("PORT", "8080"))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("app stopped")
}

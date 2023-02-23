package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/karalef/shlink/cmd"
	"github.com/karalef/shlink/repo"
	"github.com/karalef/shlink/urlservice"
	"github.com/karalef/shlink/urlservice/proto"
	"google.golang.org/grpc"
)

func init() {
	log.SetPrefix("shlink-grpc-service: ")
}

func main() {
	log.Println("starting...")
	srv := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go cmd.ListenSig(ctx, func(s os.Signal) {
		cancel()
		log.Println(s.String(), "shutting down...")
		srv.GracefulStop()
		os.Exit(0)
	})

	log.Print("connecting to mongodb...")
	mongoCfg := repo.MongoConfig{
		Addr:       cmd.Getenv("MONGO_ADDR", "mongo:27017"),
		Database:   cmd.Getenv("MONGO_DATABASE", "shlink"),
		User:       cmd.Getenv("MONGO_USER", "shlink"),
		Pass:       cmd.Getenv("MONGO_PASS", "shlink"),
		Collection: cmd.Getenv("MONGO_COLLECTION", "urls"),
	}
	rep, err := repo.NewMongoRepository(ctx, mongoCfg)
	if err != nil {
		log.Fatalln("mongodb connection failed:", err)
	}
	log.Printf("connected to mongodb://%s", mongoCfg.Addr)

	port := cmd.Getenv("PORT", "9898")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("listen on", port+":", err)
	}

	proto.RegisterURLServiceServer(srv, urlservice.NewURLService(rep))
	log.Printf("service started (listening on %s)", port)

	err = srv.Serve(lis)
	if err != nil && err != grpc.ErrServerStopped {
		log.Fatalln("failed to serve:", err)
	}

	log.Println("service stopped")
}

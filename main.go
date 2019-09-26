package main

import (
	"context"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	grpc_server "projects/TheFakeBook/internal/server"
)

const (
	grpcPort     = 9981
	gatewayServicePort = 8083
	host         = "localhost"
)

func main() {
	app := cli.NewApp()
	app.Name = "TheFakeBook"
	app.Usage = "API to manage and rate fake news"
	app.Action = launch
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func launch(_ *cli.Context) error {

	// setup mongo connection
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mongoConnectionString := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionString))
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB %v, error: %v", mongoConnectionString, err)
	}

	// start grpc server
	log.Printf("GRPC Server listenning on port %v", grpcPort)
	go grpc_server.StartService(grpc_server.StartServiceInput{
		GrpcPort:        grpcPort,
		RestServicePort: gatewayServicePort,
		MongoClient:     client,
	})

	//start REST Gateway server
	log.Printf("REST API server listenning on port %v", gatewayServicePort)
	grpc_server.StartGatewayProxy(grpc_server.GrpcGatewayParams{
		ServicePort: gatewayServicePort,
		GrpcPort:    grpcPort,
		Host:        host})

	return nil
}

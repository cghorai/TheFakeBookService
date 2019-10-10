package server

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"projects/TheFakeBook/pkg/handler"
	"projects/TheFakeBook/pkg/operation"
	"projects/TheFakeBook/pkg/service"
	"strconv"
)

type server struct {
	mongoClient    *mongo.Client
	serviceHandler *handler.FakeBookHandler
}

type StartServiceInput struct {
	MongoClient     *mongo.Client
	GrpcPort        int32
	RestServicePort int32
}

func StartService(in StartServiceInput) {

	s := grpc.NewServer()

	//Injecting Mongo Client
	repository := operation.NewFakeNewsRepository()
	repository.MongoClient = in.MongoClient
	//Injecting Repo
	service := service.NewfakeBookService()
	service.Repo = repository
	//Injecting service
	handler := &handler.FakeBookHandler{}
	handler.FakeBookService = service

	RegisterFakeBookServiceServer(s, &server{mongoClient: in.MongoClient, serviceHandler: handler})
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(9981))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) AddFakeNews(ctx context.Context, input *FakeNewsRequestWrapper) (response *FakeNewsResponse, err error) {

	id, err := s.serviceHandler.AddFakeNews(input.FakeNewsRequest.UserId, input.FakeNewsRequest.FakeNewsUrl)
	if err != nil {
		return nil, err
	}
	response = new(FakeNewsResponse)
	response.Id = id
	return response, nil
}

func (s *server) DeleteFakeNews(ctx context.Context, input *FakeNewsRequestWrapper) (*FakeNewsResponse, error) {
	panic("implement me")
}

func (s *server) RateFakeNews(ctx context.Context, input *RatingRequestWrapper) (*FakeNewsResponse, error) {
	panic("implement me")
}

func (s *server) ViewFakeNews(ctx context.Context, input *FakeNewsId) (*FakeNewsResponse, error) {
	panic("implement me")
}

func (s *server) HealthCheck(ctx context.Context, input *HealthRequest) (*HealthReply, error) {
	panic("implement me")
}

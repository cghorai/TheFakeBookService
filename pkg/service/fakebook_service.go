package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"projects/TheFakeBook/pkg/operation"
)

type FakeBookService interface {
	AddFakeNews(userId, fakeNewsUrl string) (string, error)
}

type fakeBookService struct {
	Repo operation.FakeNewsRepository
}

func NewfakeBookService() *fakeBookService {
	return &fakeBookService{}
}

func (f *fakeBookService) getFakeNewsRepository() operation.FakeNewsRepository {
	if f.Repo != nil {
		return f.Repo
	}
	f.Repo = operation.NewFakeNewsRepository()
	return f.Repo
}

func (f *fakeBookService) AddFakeNews(userId, fakeNewsUrl string) (Id string, err error) {
	if userId == "" || fakeNewsUrl == "" {
		log.Println("Blank param")
		return "", status.Error(codes.Internal, "Blank param")
	}
	id, err := f.getFakeNewsRepository().InsertFakeNews(userId, fakeNewsUrl)
	if err != nil || id == "" {
		log.Println("Couple create record for news ", fakeNewsUrl)
		return "", err
	}
	return id, nil
}

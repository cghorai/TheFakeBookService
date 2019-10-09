package handler

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"projects/TheFakeBook/pkg/service"
)

type FakeBookHandler struct {
	FakeBookService service.FakeBookService
}

func (h *FakeBookHandler) getMlsAddressService() service.FakeBookService {
	if h.FakeBookService != nil {
		return h.FakeBookService
	}
	h.FakeBookService = service.NewfakeBookService()
	return h.FakeBookService
}

func (h *FakeBookHandler) AddFakeNews(UserId string, FakeNewsUrl string) (id string, err error) {
	id, err = h.getMlsAddressService().AddFakeNews(UserId, FakeNewsUrl)
	if err != nil && err == status.Error(codes.Internal, "Database internal error") {
		return "", err
	}
	if id == "" {
		return "", errors.New("database insertion error")
	}
	//Else successfully added
	return id, err
}

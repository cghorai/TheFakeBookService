package server

import (
	"context"
	"fmt"
	_ "github.com/apex/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"log"
	"net/http"
	proto "projects/TheFakeBook/pkg/service"
)

type GrpcGatewayParams struct {
	ServicePort int
	GrpcPort    int
	Host        string
}

func StartGatewayProxy(params GrpcGatewayParams) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	_, span := trace.StartSpan(ctx, "REST Gateway")
	defer span.End()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	getMethod := fmt.Sprintf("%v:%v", params.Host, params.GrpcPort)
	err := proto.RegisterFakeBookServiceHandlerFromEndpoint(ctx, mux, getMethod, opts)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to register endpoint \n"))
	}

	addr := fmt.Sprintf(":%v", params.ServicePort)
	log.Fatal(http.ListenAndServe(addr, mux))

}

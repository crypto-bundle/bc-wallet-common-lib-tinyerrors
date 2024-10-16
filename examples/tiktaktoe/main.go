package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const Domain = "example_signer"

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrTokenExpired  = errors.New("token expired")
)

type tokenStorageService interface {
	IsTokenExistsByUUID(tokenUUID string, walletUUID string) (bool, error)
	IsTokenExpiredByUUID(tokenUUID string, walletUUID string) (bool, error)
}

type signerService interface {
	SignData(dataForSign []byte, walletUUID string) ([]byte, error)
}

type marshallerService interface {
	MarshallSignResponse(walletUUID string, signedData []byte) *SignDataResponse
	MarshallErrorResponse(err error) error
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	logger := log.Default()

	listenConn, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		logger.Fatal("unable to listen port", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	go func() {
		signHandler := &signRequestHandle{
			tokenStorageSvc: &tokenStorage{
				luckyNumber: atomic.Uint64{},
			},
			signerSvc: &dataSigner{
				luckyNumber: atomic.Uint64{},
			},
			marshallerSvc: &marshaller{},
		}

		//RegisterSignerApiServer(grpcServer, NewGrpcService(signHandler))
		grpcServer.RegisterService(&SignerApi_ServiceDesc, NewGrpcService(signHandler))

		serveErr := grpcServer.Serve(listenConn)
		if serveErr != nil {
			logger.Println("unable to start gRPC server", err)
			return
		}
	}()

	go func() {
		<-ctx.Done()

		grpcServer.GracefulStop()

		logger.Println("gRPC server successfully shutdown")
	}()

	logger.Println("application started successfully")

	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	cancelFunc()

	time.Sleep(time.Second * 5)

	logger.Println("application closed")
}

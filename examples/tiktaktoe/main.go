package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tiktaktoe/gameengine"
	"tiktaktoe/grpcserver"
	"tiktaktoe/historystore"
	"tiktaktoe/lobbyengine"
	pb "tiktaktoe/pkg"

	"github.com/crypto-bundle/bc-wallet-common-lib-tinyerrors/pkg/tinyerrors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const Domain = "example_tiktaktoe"

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	logger := log.Default()

	listenConn, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		logger.Fatal("unable to listen port", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	matchHistoryStore := historystore.NewDataStore()

	gameEngineSvc := gameengine.NewGameEngineService(tinyerrors.Default())
	lobbyEngineSvc := lobbyengine.NewLobbyEngine()

	grpcSrv := grpcserver.NewGrpcService(matchHistoryStore, gameEngineSvc, lobbyEngineSvc)

	go func() {
		//RegisterSignerApiServer(grpcServer, NewGrpcService(signHandler))
		grpcServer.RegisterService(&pb.GameApi_ServiceDesc, grpcSrv)
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

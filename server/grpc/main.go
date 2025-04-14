package grpc

import (
	"fmt"
	"net"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/client"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/services"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

type GRPCService struct {
	UsersService     pb.UserServiceServer
	FollowingService pb.FollowingServiceServer
}

func New(cfg *config.Config, log logger.Logger) (*GRPCService, neo4j.Driver, neo4j.Session, error) {
	psql, err := db.NewPostgresDB(cfg)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	grpcClient, err := client.NewGrpcClients(cfg)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error while connecting with grpc clients: %v", err)
	}

	driverNeo4j, sessionNeo4j, err := db.NewDgraphClient(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	storageObj := storage.New(psql, sessionNeo4j, log, cfg)

	serviceOptions := &services.ServiceOptions{
		ServiceManager: grpcClient,
		Storage:        storageObj,
		Config:         cfg,
		Logger:         log,
	}

	return &GRPCService{
		UsersService:     services.NewUsersService(serviceOptions),
		FollowingService: services.NewFollowUnFollowService(serviceOptions),
	}, driverNeo4j, sessionNeo4j, nil
}

func (service *GRPCService) Run(logger logger.Logger, cfg *config.Config) {
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service.UsersService)
	pb.RegisterFollowingServiceServer(server, service.FollowingService)

	listener, err := net.Listen("tcp", ":"+cfg.RPCPort)
	if err != nil {
		panic("Error while creating listener")
	}

	logger.Info("GRPC server is running on port " + cfg.RPCPort)

	err = server.Serve(listener)
	if err != nil {
		panic("error while serving gRPC server on port " + cfg.RPCPort)
	}
}

package main

import (
	"Settings/output"
	role "Settings/packages/role"
	"Settings/packages/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

func main() {
	grpcAddr := net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	// configure our core service
	userService := user.NewUserService()
	roleService := role.NewRoleService()
	// configure our gRPC service controller
	userServiceController := user.NewUserServiceController(userService)
	roleServiceController := role.NewRoleServiceController(roleService)
	// start a gRPC server
	server := grpc.NewServer()
	output.RegisterUserServiceServer(server, userServiceController)
	output.RegisterRoleServiceServer(server, roleServiceController)
	reflection.Register(server)
	con, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

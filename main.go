package main

import (
	epRole "Settings/endpoints/role"
	epUser "Settings/endpoints/user"
	"Settings/output"
	"Settings/packages/role"
	"Settings/packages/user"
	"Settings/transport"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"net/http"
	"os/signal"
	"syscall"

	"net"

	"github.com/go-kit/kit/log"
	kitGrpc "github.com/go-kit/kit/transport/grpc"
	"os"
)

const (
	defaultGRPCPort = "8082"
)

func main() {
	var (
		logger   log.Logger
		grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		userService    = user.NewUserService()
		userEndpoints  = epUser.NewEndpointSet(userService)
		userServer     = transport.NewUserServer(userEndpoints)
		roleService    = role.NewRoleService()
		roleEndpoints  = epRole.NewRoleEndpointSet(roleService)
		roleGrpcServer = transport.NewRoleServer(roleEndpoints)
	)
	var g group.Group
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			err := logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		g.Add(func() error {
			err := logger.Log("transport", "gRPC", "addr", grpcAddr)
			if err != nil {
				return err
			}
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitGrpc.Interceptor))
			output.RegisterUserServiceServer(baseServer, userServer)
			output.RegisterRoleServiceServer(baseServer, roleGrpcServer)
			grpcWebServer := grpcweb.WrapServer(
				baseServer,
				grpcweb.WithOriginFunc(func(origin string) bool { return true }),
			)
			srv := &http.Server{
				Handler: grpcWebServer,
				Addr:    fmt.Sprintf("localhost:%d", 8084),
			}
			return srv.ListenAndServe()
		}, func(error) {
			err := grpcListener.Close()
			if err != nil {
				return
			}
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	err := logger.Log("exit", g.Run())
	if err != nil {
		return
	}
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

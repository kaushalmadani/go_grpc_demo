package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "Settings/output"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:8082", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// printFeature gets the feature for the given point.
func addUser(client pb.UserServiceClient, point *pb.User) {
	log.Printf("Getting feature for point (%s, %s)", point.Name, point.Gender)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.AddUser(ctx, &pb.AddUserRequest{User: point})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}
func addRole(client pb.RoleServiceClient, point *pb.Role) {
	log.Printf("Getting feature for point (%s, %s)", point.Name, point.Status)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.AddRole(ctx, &pb.AddRoleRequest{Role: point})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = ""
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Println("closing")
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	addUser(client, &pb.User{Name: "kaushal"})
	roleClient := pb.NewRoleServiceClient(conn)
	addRole(roleClient, &pb.Role{
		Id:     1,
		Name:   "Admin",
		Status: "active",
	})
}
package main

import (
	pb "./proto"
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
)

var (
	endpoint = flag.String("echo_endpoint", "localhost:8081", "endpoint of YourService")
	user []*User
)
type User struct {
	username 	string
	password 	string
}
type server struct {}

func main(){
	go runHTTPService()
	runGRPCService()
}
func runGRPCService(){
	lis,_ := net.Listen("tcp", ":8081")
	creds, err := credentials.NewServerTLSFromFile("keys/server-cert.pem","keys/server-key.pem")
	if err!=nil{
		log.Fatalf("Failed to setup tls:%s",err)
	}
	gserver := grpc.NewServer(
			grpc.Creds(creds),
		)
	pb.RegisterAccountServer(gserver, NewServer())
	log.Fatal(gserver.Serve(lis))

}
func NewServer() *server {
	return &server{}
}
func runHTTPService(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile("keys/server-cert.pem","")
	if err != nil{
		log.Fatalf("gateway cert load error: %s\n",err)
		return
	}
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	if err:=pb.RegisterAccountHandlerFromEndpoint(ctx,mux,*endpoint,opts);err!=nil{
		log.Fatalf("failed to start HTTP server: %v", err)
	}
	log.Printf("Http is listening on port:8080")
	log.Fatal(http.ListenAndServe(":8080",mux))
}

func (s *server) Create(ctx context.Context, in *pb.Request) (*pb.Response, error){
	username := in.GetUsername()
	if username == ""{
		return nil, status.New(codes.InvalidArgument,"username could not be empty").Err()
	}

	passwd := in.GetPassword()
	if passwd == ""{
		return nil,status.New(codes.InvalidArgument,"password could not be empty").Err()
	}

	u := new(User)
	u.username = username
	u.password = passwd
	user = append(user,u)
	return &pb.Response{
		Status:"success",
	},nil
}

func (s *server) Login(ctx context.Context, in *pb.Request) (*pb.Response, error){
	token := ctx.Value("x-token")
	fmt.Println(token)
	return nil,nil
}
func (s *server) GetAccount(ctx context.Context, in *pb.Request) (*pb.Response, error){
	return nil,nil
}
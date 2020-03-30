package main
import(
	"net/http"
	"context"
	// "github.com/ti/noframe/grpcmux"
	"log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"fmt"
	"flag"
	"net"
	pb "./pb"
)

type server struct{}
// type ServerClient struct {
// 	srv SayServer
// }
// func RegisterSayServerHandlerClient(ctx context.Context, mux *runtime.ServeMux, srv SayServer) error {
// 	return RegisterLoginHandlerClien(ctx, mux, NewLoginClient(srv))
// }
var (
    echoEndpoint = flag.String("echo_endpoint", "localhost:8081", "endpoint of YourService")
)

func main(){
	srv := &server{}

	go func(){//http server 
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		pb.RegisterLoginHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
		http.ListenAndServe(":8080", mux)
	}()
	//serve Grpc call
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	gs := grpc.NewServer()
	pb.RegisterLoginServer(gs, srv)
	if err := gs.Serve(lis); err!=nil{
		log.Fatalf("failed to serve: %v", err)
	}
}
func (s *server) Hello(ctx context.Context, request *pb.LoginInfo) (*pb.Response, error){
	fmt.Println(request.Username)
	fmt.Println(request.Password)
	return &pb.Response{
		Status:  "success",
	}, nil	
}
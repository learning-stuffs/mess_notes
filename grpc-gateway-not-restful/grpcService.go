package main
import (
  "context"  // Use "golang.org/x/net/context" for Golang version <= 1.6
  // "flag"
  // "net/http"
  // "github.com/golang/glog"
  // "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "net"
  "google.golang.org/grpc"

  pb "./src/proto"  // Update
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
    return &pb.StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
    return &pb.StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}

func (r *RestServiceImpl) SayHello(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value:"Hi "+ message.Value + "from sayHello"},nil
}
func main() {
    grpcServer := grpc.NewServer()
    pb.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
    lis, _ := net.Listen("tcp", ":5000")
    grpcServer.Serve(lis)
}
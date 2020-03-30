package main

import (
  "context"  // Use "golang.org/x/net/context" for Golang version <= 1.6
  "net/http"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
  "log"
  gw "./src/proto"  // Update
)


func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()

    err := gw.RegisterRestServiceHandlerFromEndpoint(
        ctx, mux, "localhost:5000",
        []grpc.DialOption{grpc.WithInsecure()},
    )
    if err != nil {
        log.Fatal(err)
    }

    http.ListenAndServe(":8080", mux)
}
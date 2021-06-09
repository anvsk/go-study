package main

import (
    "io"
    "log"
    "net/http"
    "net/rpc"
    "net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello:" + request
    return nil
}

// const HelloServiceName = "path/to/pkg.HelloService"
const HelloServiceName = "HelloService"

type HelloServiceInterface interface {
    Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
    return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
    *rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
    return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
    server()
    // client()
}

func server() {
    rpc.RegisterName(HelloServiceName, new(HelloService))

    http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
        var conn io.ReadWriteCloser = struct {
            io.Writer
            io.ReadCloser
        }{
            ReadCloser: r.Body,
            Writer:     w,
        }

        rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
    })

    http.ListenAndServe(":1234", nil)
}

func client() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    var reply string
    err = client.Call(HelloServiceName+".Hello", "hello", &reply)
    if err != nil {
        log.Fatal(err)
    }
}

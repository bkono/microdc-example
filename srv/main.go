package main

import (
	"fmt"
	"log"
	"time"

	hello "github.com/bkono/microdc-example/srv/proto/hello"
	vip "github.com/bkono/microdc-example/vip-srv/proto/vip"
	"github.com/micro/go-micro"

	"golang.org/x/net/context"
)

var (
	regGreet = "Hello %s"
	vipGreet = "Well hello, %s. Thanks for being a VIP!"
)

type Say struct {
	vipcl vip.VIPClient
}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request, checking vip")

	viprsp, err := s.vipcl.CheckVIP(ctx, &vip.VIPRequest{Name: req.Name})
	log.Println(viprsp, err)

	if viprsp.IsVip {
		rsp.Msg = fmt.Sprintf(vipGreet, req.Name)
	} else {
		rsp.Msg = fmt.Sprintf(regGreet, req.Name)
	}

	return nil
}

func NewSayHandler(client vip.VIPClient) hello.SayHandler {
	return &Say{client}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Setup clients
	cl := vip.NewVIPClient("go.micro.srv.vip", service.Client())

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), NewSayHandler(cl))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

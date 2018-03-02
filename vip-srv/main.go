package main

import (
	"github.com/bkono/microdc-example/vip-srv/handler"
	pb "github.com/bkono/microdc-example/vip-srv/proto/vip"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.vip"),
		micro.Version("latest"),
	)

	// Register Handler
	pb.RegisterVIPHandler(service.Server(), handler.NewVIPHandler())

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

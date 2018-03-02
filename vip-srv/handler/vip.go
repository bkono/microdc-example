package handler

import (
	"log"
	"math/rand"

	pb "github.com/bkono/microdc-example/vip-srv/proto/vip"
	"golang.org/x/net/context"
)

type vipHandler struct{}

func (v *vipHandler) CheckVIP(ctx context.Context, req *pb.VIPRequest, rsp *pb.VIPResponse) error {
	log.Print("Received CheckVIP request, checking vip")

	rsp.IsVip = rand.Intn(10) > 5
	log.Println("is vip check", rsp.IsVip)

	return nil
}

func NewVIPHandler() pb.VIPHandler {
	return &vipHandler{}
}

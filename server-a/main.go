package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"gomod.alauda.cn/integ-grpc/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*product.Product
}

//添加商品
func (s *server) AddProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	req.Name = fmt.Sprintf("%s-A", req.Name)
	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

//获取商品
func (s *server) GetProduct(ctx context.Context, req *product.ProductId) (resp *product.Product, err error) {
	fmt.Println("get project from A,", req.Value)
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	resp = s.productMap[req.Value]
	return
}

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	product.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

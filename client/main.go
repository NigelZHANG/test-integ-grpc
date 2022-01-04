package main

import (
	"context"
	"fmt"
	"log"

	"gomod.alauda.cn/integ-grpc/product"
	"google.golang.org/grpc"
)

const (
	addressA = "localhost:8080"
	addressB = "localhost:8081"
)

func main() {
	ctx := context.Background()
	// client A
	connA, err := grpc.Dial(addressA, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer connA.Close()

	clientA := product.NewProductInfoClient(connA)

	id := AddProduct(ctx, clientA)
	GetProduct(ctx, clientA, id)
	fmt.Println("--------------")

	// client B
	connB, err := grpc.Dial(addressB, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer connB.Close()

	clientB := product.NewProductInfoClient(connB)

	idB := AddProduct(ctx, clientB)
	GetProduct(ctx, clientB, idB)
}

// 添加一个测试的商品
func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	aMac := &product.Product{Name: "name", Description: "test"}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success, id = ", productId.Value)
	return productId.Value
}

// 获取一个商品
func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	p, err := client.GetProduct(ctx, &product.ProductId{Value: id})
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Printf("get prodcut success : %+v\n", p)
}

package main


import (
	"awesomeProject/grpc/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var (
	port = flag.Int("port",9000,"the port connect")
)
type server struct {
	proto.UnimplementedOrderServiceServer
}

func (s * server) NewOrder(ctx context.Context,in *proto.NewRequestOrder)(*proto.NewResponseOrder,error){
	log.Printf("order id: %d %d",in.GetOrderRequest())
	orderId := in.GetOrderRequest()  // Lấy orderId từ yêu cầu của client

	// Giả sử bạn lấy dữ liệu từ một cơ sở dữ liệu hoặc một nguồn khác
	orderName := "Sample Order"

	intValue, err := strconv.ParseInt(in.GetOrderRequest(), 0, 32)
	if err  != nil{
		fmt.Println("loi convet ")
	}
	return  &proto.NewResponseOrder{
		OrderResponse: "newOrderId" + fmt.Sprintf("%d", orderId),
		OrderName:     orderName,
		OrderId: int32(intValue),
		OrderPayment:  10000,
	}, nil
}

func main() {
	lis,err := net.Listen("tcp",fmt.Sprintf(":%d",*port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, &server{})
	log.Printf("sever adddress %s",lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

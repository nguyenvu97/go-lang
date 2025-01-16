package main

import (
	"awesomeProject/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	address := "localhost:9000"

	// Kết nối đến server gRPC
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Khởi tạo client
	c := proto.NewOrderServiceClient(conn)

	// Tạo ticker mỗi 2 giây
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		orderId := "10001"

		// Tạo context với timeout 1 giây
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel() // Đảm bảo hủy context sau khi sử dụng

		// Gọi gRPC method
		r, err := c.NewOrder(ctx, &proto.NewRequestOrder{
			OrderRequest: orderId,
		})

		// Kiểm tra lỗi khi gọi method
		if err != nil {
			log.Printf("Error while creating order: %v", err)
			continue // Tiếp tục vòng lặp nếu có lỗi
		}

		// Log kết quả trả về
		log.Printf("Order Response: %s %s %s", r.GetOrderResponse(),r.GetOrderName(),r.GetOrderPayment())
	}
}

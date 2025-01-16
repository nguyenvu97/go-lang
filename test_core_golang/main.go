package main

import (
	"awesomeProject/test_core_golang/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Xử lý yêu cầu đến server
func handleRequest(c *gin.Context) {
	// Giả lập xử lý yêu cầu trong goroutine
	go func() {
		time.Sleep(2 * time.Second) // Giả lập công việc xử lý tốn thời gian
		fmt.Println("Request processed in goroutine")
	}()

	c.JSON(200, gin.H{
		"message": "Request received and processed in goroutine",
	})
}

// Gửi yêu cầu đến server
func sendRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("http://localhost:1122/ticket/1/detail/1")
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var data  dto.ResponseWithInterface
	err = json.Unmarshal(body, &data) // Dùng json.Unmarshal để chuyển đổi
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	if data.Success {


		switch v := data.Result.(type) {
		case map[string]interface{}:
			var ticketDto dto.TicketDto
			ticketBytes, err := json.Marshal(v)
			if err != nil {
				fmt.Println("Error unmarshalling result to TicketDto:", err)
				return
			}
			err = json.Unmarshal(ticketBytes, &ticketDto)
			fmt.Printf("Parsed TicketDto: %+v\n", ticketDto)
		default:
			fmt.Println("Unexpected result type:", v)
		}
	}else {
		fmt.Println("Response was not successful.")
	}


}

func main() {
	// Khởi chạy server trong goroutine riêng
	//go func() {
	//	r := gin.Default()
	//	r.GET("/process", handleRequest)
	//	_ = r.Run(":8080") // Server chạy trên localhost:8080
	//}()

	// Đợi server khởi động
	time.Sleep(1 * time.Second)
	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	var wg sync.WaitGroup
	numRequests := 100
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go sendRequest(&wg)
	}

	wg.Wait() // Chờ tất cả các yêu cầu hoàn thành
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("time data : %d",endTime - startTime)
	fmt.Println("All requests completed")
}


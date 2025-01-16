package main

import (
	"fmt"
	"modulB/src/dto"
	"time"
)

//func main() {
//	var wg sync.WaitGroup
//
//	numRequests := 2
//	for i := 0; i < numRequests; i++ {
//		wg.Add(1)
//		sendRequest(&wg)
//	}
//		wg.Wait()
//}
//func Sub(channel <- chan dto.TicketDto , userName string,wg *sync.WaitGroup)  {
//	defer wg.Done()
//	for msg := range channel{
//		fmt.Printf("username %s :: data %+v\n", userName, msg)
//		time.Sleep(time.Second *1)
//	}
//
//
//}
//func Pub(channel chan <- dto.TicketDto , data dto.TicketDto , number int){
//	for i := 1 ; i <= number; i++{
//		channel <- data
//		fmt.Printf("Publisher gửi data: %+v\n", data)
//		time.Sleep(time.Second * 10) // Giả lập thời gian gửi
//	}
//
//
//	close(channel) // Đóng channel sau khi gửi xong
//}
//
//func sendRequest(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	resp, err := http.Get("http://localhost:1122/ticket/1/detail/1")
//	if err != nil {
//		fmt.Println("Error sending request:", err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Error reading response body:", err)
//		return
//	}
//	var data  dto.Response
//	err = json.Unmarshal(body, &data) // Dùng json.Unmarshal để chuyển đổi
//	if err != nil {
//		fmt.Println("Error unmarshalling JSON:", err)
//		return
//	}
//	if data.Success {
//		ch := make(chan dto.TicketDto)
//		var ticketDto dto.TicketDto
//		err = json.Unmarshal(data.Result, &ticketDto) // Giải mã Result
//		if err != nil {
//			fmt.Println("Error unmarshalling result to TicketDto:", err)
//			return
//		}
//		numSubscribers := 3
//		go Pub(ch,ticketDto,numSubscribers)
//
//		for i := 1; i <= numSubscribers; i++ {
//			wg.Add(1)
//
//			go Sub(ch,"nguyenvu",wg)
//		}
//		wg.Wait()
//	}
//
//
//}

func main() {
	buyCh := make(chan dto.TicketDto)
	cancelCh := make(chan string)

	cancelTicket := []string{"1","2"}


	var dataTicket = dto.Testdata()
	go BuyTicket(buyCh,dataTicket)
	go CancelTicket(cancelCh,cancelTicket)
	go HandlerTicketWithSelect(buyCh,cancelCh)
	time.Sleep(time.Second *30)
	fmt.Print("end buy and canceling ...")


}
func BuyTicket(channel chan <- dto.TicketDto, data []dto.TicketDto){
	for _,order := range  data{
		time.Sleep(time.Second * 2)
		channel <- order // send in channel
	}
	close(channel)
}
func CancelTicket(channel chan <- string, data []string)  {
	for _,orderId := range  data{
		time.Sleep(time.Second * 10)
		channel <- orderId // send in channel

	}
	close(channel)
	
}
func HandlerTicket(orderChannel <- chan dto.TicketDto, cancelTicket <- chan string)  {

	for  {
		orderBuyTicketOk, err := <- orderChannel // read gorouting data
		if err {
			fmt.Printf(" Handler buy ticket ... %s  - %s - %d\n",orderBuyTicketOk.Id,orderBuyTicketOk.Name,orderBuyTicketOk.Description)
		}else {
			fmt.Println("order channel is closed")
			break
		}
		cancelTicketOk, cancelTicket := <- cancelTicket // read gorouting data
		if cancelTicket {
			fmt.Printf("Handler cancel ticket ... %d\n",cancelTicketOk)
		}else {
			fmt.Println("cancel channel is closed")
			break
		}
	}


}
func HandlerTicketWithSelect(orderChannel <- chan dto.TicketDto, cancelChannel <- chan string)  {

	for  {
		select {
			case orderBuyTicketOk, err := <- orderChannel : // read gorouting data
			if err {
				fmt.Printf(" Handler buy ticket ... %s  - %s - %d\n",orderBuyTicketOk.Id,orderBuyTicketOk.Name,orderBuyTicketOk.Description)
			}else {
				fmt.Println("order channel is closed")
				orderChannel = nil
			}
			case cancelTicketOk, cancelTicket := <- cancelChannel : // read gorouting data
			if cancelTicket {
				fmt.Printf("Handler cancel ticket ... %d\n",cancelTicketOk)
			}else {
				fmt.Println("cancel channel is closed")
				cancelChannel = nil
			}
		}
		if orderChannel == nil && cancelChannel == nil {
			break
		}

	}


}
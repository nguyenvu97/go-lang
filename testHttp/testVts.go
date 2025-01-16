package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

//func main() {
//
//	numTasks := 100000 // Số lượng tác vụ cần xử lý
//	var wg sync.WaitGroup
//	stratTime := time.Now().UnixNano() / int64(time.Millisecond)
//	// Tạo các goroutines để xử lý tác vụ
//	for i := 0; i < numTasks; i++ {
//		wg.Add(1)
//		go processTask(&wg, i)
//	}
//
//	// Chờ cho tất cả goroutines hoàn thành
//	wg.Wait()
//	endTime := time.Now().UnixNano() / int64(time.Millisecond)
//	// In thông báo khi tất cả tác vụ đã hoàn thành
//	fmt.Printf("time data : %d",endTime - stratTime)
//	fmt.Println("All tasks are processed.")
//}
//func processTask (wg *sync.WaitGroup, taskID int)  {
//	defer wg.Done()
//
//	// In ra tên goroutine đang xử lý
//	fmt.Printf("Processing task %d in goroutine: %s\n", taskID, fmt.Sprintf("goroutine-%d", taskID))
//
//	// Giả lập công việc mất thời gian
//	time.Sleep(100 * time.Millisecond)
//}
type MyError struct {
	When time.Time
	Where string
}

func (e * MyError) Error()string  {
	return  fmt.Sprintf("at %v, %s",
		e.When, e.Where)
}

func Run() error {
	return &MyError{
		When: time.Now(),
		Where: "service",
	}
}
type rot13Reader struct {
	r io.Reader
}


func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Hàm kiểm tra mật khẩu
func checkPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func main() {
	// Tạo hash của mật khẩu "hello"
	//hashedPassword, err := hashPassword("nguyenvu")
	//if err != nil {
	//	fmt.Println("Error hashing password:", err)
	//	return
	//}
	//var name string
	//
	//fmt.Print("Nhập tên của bạn: ")
	//fmt.Scanln(&name)
	//
	// err1 := checkPassword(name,hashedPassword)
	//if err1 != nil {
	//	fmt.Printf("Error hashing password:")
	//	return
	//}
	//fmt.Println("login Ok")
	hashedPassword := "$2a$10$QZxvFyl7t0pZJo1XrhRLsu2ohl9F53g9MHKEbPakhweS1EYZG7oxy"

	// Mật khẩu mà bạn muốn kiểm tra
	inputPassword := ""

	// So sánh mật khẩu thô với hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		fmt.Println("Mật khẩu không khớp!")
	} else {
		fmt.Println("Mật khẩu khớp!")
	}





}
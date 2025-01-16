package main


import "fmt"

func getPersonInfo() (string, int, bool) {
	name := "Alice"
	age := 30
	isStudent := false
	return name, age, isStudent
}
func getPersonInf1() (string, int, bool) {
	name, age, isStudent := getPersonInfo()
	name = "nguyenvu"
	age = 30
	isStudent = false
	return name, age, isStudent
}
func main() {
	name, age ,isStudent := getPersonInf1()
	fmt.Printf("name %d  -- age : %a  --- isStudent : %d",name,age,isStudent )
}
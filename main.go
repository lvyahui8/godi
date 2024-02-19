package main

func main() {
	service := GetBean[*UserService]()
	service.SayHello()
}

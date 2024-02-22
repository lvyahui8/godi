package main

func main() {
	service := GetBean[*NormalUserService]()
	service.SayHello()

	controller := GetBean[*HomeController]()
	controller.Index()
}

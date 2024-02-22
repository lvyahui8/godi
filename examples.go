package main

import "fmt"

type UserRepo struct {
	*Bean
}

func (repo *UserRepo) GetId() int {
	return 1
}

type UserDao struct {
	*Bean
	Repo *UserRepo
}

func (dao *UserDao) GetName() string {
	return fmt.Sprintf("sam_%02d", dao.Repo.GetId())
}

type UserService interface {
	SayHello()
}

type NormalUserService struct {
	*Bean
	Dao *UserDao
}

func (us *NormalUserService) SayHello() {
	fmt.Println("hello " + us.Dao.GetName())
}

type KeyUserService struct {
	*Bean
}

func (us *KeyUserService) SayHello() {
	fmt.Println("hello keyUser")
}

type HomeController struct {
	*Bean
	UserService *UserService `inject:"KeyUserService"`
	content     string
}

func (hc *HomeController) Init() error {
	hc.content = "sam"
	return nil
}

func (hc *HomeController) Index() {
	fmt.Println("hello " + hc.content)
}

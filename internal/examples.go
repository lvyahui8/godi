package internal

import (
	"fmt"
	"github.com/lvyahui8/godi"
)

type UserRepo struct {
	*godi.Bean
}

func (repo *UserRepo) GetId() int {
	return 1
}

type UserDao struct {
	*godi.Bean
	Repo *UserRepo
}

func (dao *UserDao) GetName() string {
	return fmt.Sprintf("sam_%02d", dao.Repo.GetId())
}

type UserService interface {
	SayHello()
}

type NormalUserService struct {
	*godi.Bean
	Dao *UserDao
}

func (us *NormalUserService) SayHello() {
	fmt.Println("hello " + us.Dao.GetName())
}

type KeyUserService struct {
	*godi.Bean
}

func (us *KeyUserService) SayHello() {
	fmt.Println("hello keyUser")
}

type HomeController struct {
	*godi.Bean
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

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

type UserService struct {
	*Bean
	Dao *UserDao
}

func (us *UserService) SayHello() {
	fmt.Println("hello " + us.Dao.GetName())
}

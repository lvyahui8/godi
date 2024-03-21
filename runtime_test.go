package main

import "testing"

type HomeApi interface {
	Index()
}

func TestGetBeansOfType(t *testing.T) {
	GetBeansOfType[HomeApi]()
}

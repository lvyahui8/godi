package main

import "reflect"

var beanType = reflect.TypeOf((*BeanSupport)(nil)).Elem()
var beanTypeStr = beanType.String() //  "*main.Bean"

func IsBean(v any) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Struct {
		t = reflect.PointerTo(t)
	}
	// Implements方法会遍历比对方法来判断是否实现了interface
	return t.Implements(beanType)
}

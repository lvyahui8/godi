package main

import (
	"reflect"
)

type BeanSupport interface {
	BeanName() string
	Init()
}

type Bean struct {
}

func (b *Bean) BeanName() string {
	return ""
}

func (b *Bean) Init() {

}

var beanMap = make(map[string]any)

var beanType = reflect.TypeOf((*BeanSupport)(nil)).Elem()

func GetBean[T BeanSupport]() T {
	var t T
	return getBean(t).(T)
}

func getBean(v any) any {
	var t BeanSupport
	t = v.(BeanSupport)
	var beanName string
	val := reflect.ValueOf(t)
	if !val.IsNil() {
		beanName = t.BeanName()
	}
	if beanName == "" {
		beanName = val.Type().String()
	}
	if b, ok := beanMap[beanName]; ok {
		return b
	}
	val = reflect.New(reflect.TypeOf(t).Elem())
	// 懒加载构造和初始化bean
	// 遍历属性，看是否需要依赖注入
	elem := val.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		kind := field.Kind()
		fType := field.Type()
		if fType.String() == "*main.Bean" {
			continue
		}
		if kind != reflect.Pointer && kind != reflect.Struct {
			continue
		}

		if kind == reflect.Pointer {
			val := field.Interface()
			fType = reflect.TypeOf(val)
		}
		if fType.Implements(beanType) {
			// 递归调用getBean，组装bean
			field.Set(reflect.ValueOf(getBean(field.Interface())))
		}
	}
	t = val.Interface().(BeanSupport)
	t.Init()
	beanMap[beanName] = val.Interface()
	return val.Interface()
}

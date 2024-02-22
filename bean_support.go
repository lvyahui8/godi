package main

import (
	"reflect"
)

// BeanSupport bean的基类型
type BeanSupport interface {
	// BeanName bean可以自定义名称。默认名称为全类型限定名
	BeanName() string
	// Init 支持Bean在依赖注入完成后执行Init操作
	Init() error
}

// Bean 用于嵌套在struct中，声明一个Bean
type Bean struct {
}

func (b *Bean) BeanName() string {
	return ""
}

func (b *Bean) Init() error {
	return nil
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
	err := t.Init()
	if err != nil {
		panic(err)
	}
	beanMap[beanName] = val.Interface()
	return val.Interface()
}

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

// BeanName Bean可以重写这个方法以定义bean名称。
func (b *Bean) BeanName() string {
	return ""
}

// Init Bean可以重写这个方法以完成一些依赖注入完成之后的初始化动作。
func (b *Bean) Init() error {
	return nil
}

func selfInject(v any) any {
	var t BeanSupport
	t = v.(BeanSupport)
	val := reflect.New(reflect.TypeOf(t).Elem())
	// 懒加载构造和初始化bean
	// 遍历属性，看是否需要依赖注入
	elem := val.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		kind := field.Kind()
		fType := field.Type()
		if fType.String() == beanTypeStr {
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
			//field.Set(reflect.ValueOf(getBean(field.Interface())))
		}
	}
	t = val.Interface().(BeanSupport)
	err := t.Init()
	if err != nil {
		panic(err)
	}

	return val.Interface()
}

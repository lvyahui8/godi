package main

import "reflect"

var beanType = reflect.TypeOf((*BeanSupport)(nil)).Elem()

func GetBeanByName[T BeanSupport](name string) (t T, b bool) {
	return
}

// GetBean 根据泛型参数类型获取Bean
func GetBean[T BeanSupport]() (T, bool) {
	var t T
	//return getBean(t).(T)
	return t, false
}

func IsBean(v any) bool {
	return reflect.TypeOf(v).Implements(beanType)
}

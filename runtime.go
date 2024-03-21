package main

func GetBeanByName(name string) (bean any, b bool) {
	return
}

// GetBean 根据泛型参数类型获取Bean
func GetBean[T any]() (T, bool) {
	var t T
	return t, false
}

// GetBeansOfType 根据interface获取一组bean
func GetBeansOfType[T any]() ([]any, bool) {
	//var t T
	//if reflect.TypeOf(t).Kind() == reflect.Interface {
	//	// 泛型类型参数必须是一个interface
	//	return nil, false
	//}
	return nil, true
}

// GetAllBeans 获取全部的Bean
func GetAllBeans() []any {
	return nil
}

// GetAllBeanNames 	获取全部的Bean的名称
func GetAllBeanNames() []string {
	return nil
}

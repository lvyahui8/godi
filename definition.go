package main

import "reflect"

// BeanProps bean的定制化属性
type BeanProps struct {
	// Name Bean名称
	Name string
	// Private 是否私有
	Private bool
	// Init Bean注入完依赖之后被自动调用的初始化方法。如果Bean本身实现了BeanInitializer接口，则这个方法不会调用
	Init PostConstruct
}

// beanDefinition bean的定义
type beanDefinition struct {
	BeanProps
	// beanType bean的具体类型
	beanType reflect.Type
	// implements bean实现的所有接口
	implements []reflect.Type
}

// beanInstance bean实例
type beanInstance struct {
	*beanDefinition
	// object 实际的bean对象
	object any
	// created 实际对象是否由框架创建
	created bool
}

// Bean 用于嵌套在struct中，声明一个Bean
type Bean struct {
}

// BeanCustomizer 支持通过实现方法来自定义bean
type BeanCustomizer interface {
	// GetBeanProps 支持BeanProps来定制化bean
	GetBeanProps() BeanProps
}

// BeanInitializer 支持通过实现方法来做Bean依赖注入完成后的初始化工作
type BeanInitializer interface {
	// Init 支持Bean初始化，在Bean的依赖被注入完成后自动调用，如果初始化失败，则退出容器
	Init() error
}

type PostConstruct func(bean any) error

// Register 支持往容器内手工注入一个Bean
func Register(obj any, props BeanProps) {

}

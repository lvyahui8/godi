package godi

import "reflect"

type beanName string
type typeName string

// BeanProps bean的定制化属性
type BeanProps struct {
	// Name Bean名称
	Name beanName
	// AlwaysNew 默认是否重复使用现有实例，还是每次构建最新的
	// 并不表示严格单例，某些情况，即使AlwaysNew=false, 也可能会new新的，比如tag为private私有注入时
	AlwaysNew bool
	// Init Bean注入完依赖之后被自动调用的初始化方法。如果Bean本身实现了BeanInitializer接口，则这个方法不会调用
	Init PostConstruct
}

// beanDefinition bean的定义
type beanDefinition struct {
	BeanProps
	// beanType bean的具体类型
	beanType reflect.Type
	// bean类型对应的简写
	tName typeName
}

// beanInstance bean实例
type beanInstance struct {
	*beanDefinition
	// object 实际的bean对象. 如果bean是private的，则这个值永远为空，每次获取这个依赖都构造新的对象
	object any
	// reflectValue 缓存object的反射表示
	reflectValue reflect.Value
	// created 实际对象是否由框架创建
	created bool
	// completed bean是否已经构造完成
	completed bool
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

func newBeanInstance(obj any, autoCreated bool, props BeanProps) *beanInstance {
	beanType := reflect.TypeOf(obj)
	// 构造beanInstance，并放入graph中
	definition := &beanDefinition{
		BeanProps: props,
		beanType:  beanType,
		tName:     getTypeName(beanType),
	}
	if definition.Name == "" {
		// 未指定bean名称，则默认用类型名称作为bean名称
		definition.Name = beanName(definition.tName)
	}

	// 构造新的bean实例，放入g容器
	return &beanInstance{object: obj, reflectValue: reflect.ValueOf(obj), beanDefinition: definition, created: autoCreated}
}

// Register 支持往容器内手工注入一个Bean
func Register(obj any, props BeanProps) *DIError {
	instance := newBeanInstance(obj, false, props)
	if old, exist := g.nodes[instance.Name]; exist {
		// 不允许同名的bean
		return ErrSameBeanName.CreateError(nil, instance.Name, old.instance.tName, instance.tName)
	}
	g.addNodeByInstance(instance)
	return nil
}

func getTypeName(t reflect.Type) typeName {
	return typeName(t.String())
}

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

// BeanName Bean可以重写这个方法以定义bean名称。
func (b *Bean) BeanName() string {
	return ""
}

// Init Bean可以重写这个方法以完成一些依赖注入完成之后的初始化动作。
func (b *Bean) Init() error {
	return nil
}

type BeanKey string

type Node struct {
	key       BeanKey
	object    any
	completed bool
	private   bool
	In        map[BeanKey]*Edge
	Out       map[BeanKey]*Edge
}

type Edge struct {
	from BeanKey
	to   BeanKey
	// 指针依赖还是struct依赖
	ptr bool
}

// Graph bean的依赖关系图
type Graph struct {
	nodes map[BeanKey]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[BeanKey]*Node),
	}
}

func (g *Graph) AddEdge(from, to Node) {
	if from.Out != nil {
		if _, ok := from.Out[to.key]; ok {
			// 边已经存在
			return
		}
	}
	edge := &Edge{from: from.key, to: to.key}
	from.Out[to.key] = edge
	to.In[from.key] = edge
}

func getBeanKey(t BeanSupport) BeanKey {
	var beanName string
	val := reflect.ValueOf(t)
	if !val.IsNil() {
		beanName = t.BeanName()
	}
	if beanName == "" {
		beanName = val.Type().String()
	}
	return BeanKey(beanName)
}

func inject(v any) any {
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

	return val.Interface()
}

func getBean(v any) any {
	beanName := getBeanKey(v.(BeanSupport))
	if b, ok := container.GetBean(beanName); ok {
		return b
	}

	inject(v)
	//container.PutBean(beanName, val.Interface())
	return nil
}

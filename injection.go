package main

import (
	"github.com/fatih/structtag"
	"reflect"
)

// TagAutowired tag key
const TagAutowired = "autowired"

// AutowiredTags 用来声明成员需要进行依赖注入
// eg.
// autowired:"xxxService,private"
// autowired:"private"
// autowired:"private,optional"
// autowired:"optional,private"
// autowired:"optional"
// autowired:"xxxService,optional"
type AutowiredTags struct {
	BeanProps
	// Optional 是否可选注入，默认为必须注入
	Optional bool
}

func isNotZeroValue(v reflect.Value, t reflect.Type) bool {
	switch v.Kind() {
	default:
		return !reflect.DeepEqual(v.Interface(), reflect.Zero(t).Interface())
	case reflect.Interface, reflect.Ptr:
		return !v.IsNil()
	}
}

func parseTag(fieldDef reflect.StructField) (*AutowiredTags, *DIError) {
	tags, err := structtag.Parse(string(fieldDef.Tag))
	if err != nil {
		return nil, ErrTagParseFailed.CreateError(&err, fieldDef.Name, fieldDef.Tag)
	}
	tag, err := tags.Get(TagAutowired)
	if err != nil { // error 必然为errTagNotExist
		return nil, nil
	}
	res := &AutowiredTags{}
	if tag.Name == "private" {
		res.Private = true
	} else if tag.Name == "optional" {
		res.Optional = true
	} else {
		res.Name = beanName(tag.Name)
	}
	for _, option := range tag.Options {
		if option == "private" {
			res.Private = true
		} else if option == "optional" {
			res.Optional = true
		}
	}
	return res, nil
}

// AutoInject 手动触发完整图的依赖注入
func AutoInject() *DIError {
	// 对依赖图进行动态遍历（会不断有新的成员加入到图中）
	var roots []beanName
	// 迭代的过程中，g.nodes会不断有新的元素加入（一定是非root节点），所以分两次，先取出root
	for name, node := range g.nodes {
		if len(node.edgesIn) == 0 {
			roots = append(roots, name)
		}
	}
	for _, name := range roots {
		diError := inject(g.nodes[name])
		if diError != nil {
			return diError
		}
	}
	return nil
}

// inject 递归广度优先遍历依赖图，完成图中节点的依赖注入
func inject(node *beanNode) *DIError {
	// 遍历当前bean的所有成员字段，看哪些字段需要依赖注入
	beanType := node.instance.beanType
	elem := node.instance.reflectValue.Elem()
	for i := 0; i < elem.NumField(); i++ {
		fieldValue := elem.Field(i)
		fType := fieldValue.Type()
		if !fieldValue.CanSet() {
			continue
		}
		// fieldValue.IsZero() 可能会抛panic
		if isNotZeroValue(fieldValue, fType) {
			// 对于指针、接口、map/slice类型，已经有值的成员不覆盖
			continue
		}
		// 解析autowired tag
		fieldDef := beanType.Elem().Field(i)
		tag, diError := parseTag(fieldDef)
		if diError != nil {
			return diError
		}
		if tag == nil {
			continue
		}
		// 根据tag决策注入逻辑。
	}
	// 继续广度遍历子节点
	for dependName := range node.edgesOut {
		diError := inject(g.nodes[dependName])
		if diError != nil {
			return diError
		}
	}
	return nil
}

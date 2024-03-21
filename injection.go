package main

// TagAutowired tag key
const TagAutowired = "autowired"

// AutowiredTags 用来声明成员需要进行依赖注入
type AutowiredTags struct {
	BeanProps
	// Optional 是否可选注入，默认为必须注入
	Optional bool
}

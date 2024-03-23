package main

var g *DependGraph = &DependGraph{
	nodes: make(map[beanName]*beanNode),
}

type depend struct {
	// consumer from 消费方，使用方
	consumer beanName
	// provider to 提供方，被注入的依赖
	provider beanName
}

type beanNode struct {
	instance *beanInstance
	// edgesOut key为当前bean所依赖的bean名称
	edgesOut map[beanName]*depend
	// edgesIn key为依赖当前bean的bean名称
	edgesIn map[beanName]*depend
}

// DependGraph 依赖图。 构建依赖关系过程中，往依赖图中添加依赖
type DependGraph struct {
	nodes map[beanName]*beanNode
}

func (dg *DependGraph) addNode(instance *beanInstance) {
	if _, exist := dg.nodes[instance.Name]; exist {
		return
	}
	dg.nodes[instance.Name] = &beanNode{
		instance: instance,
		edgesOut: make(map[beanName]*depend),
		edgesIn:  make(map[beanName]*depend),
	}
}

func (dg *DependGraph) addEdge(from, to beanName) {
	fromNode := dg.nodes[from]
	toNode := dg.nodes[to]
	d := &depend{consumer: from, provider: to}
	fromNode.edgesOut[to] = d
	toNode.edgesIn[from] = d
}

func (dg *DependGraph) addNodeEdge(fromNode, toNode *beanNode) {
	from := fromNode.instance.Name
	to := fromNode.instance.Name
	d := &depend{consumer: from, provider: to}
	fromNode.edgesOut[to] = d
	toNode.edgesIn[from] = d
}

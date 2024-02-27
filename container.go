package main

var container = &BeanContainer{g: NewGraph()}

type BeanContainer struct {
	g *Graph
}

func (bc *BeanContainer) PutBean(name BeanKey, node *Node) {
	bc.g.nodes[name] = node
}

func (bc *BeanContainer) GetBean(name BeanKey) (*Node, bool) {
	n, ok := bc.g.nodes[name]
	return n, ok
}

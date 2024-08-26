package ast

type Visitor interface {
	visit(node Node) any
}

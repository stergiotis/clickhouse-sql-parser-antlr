package ast

import (
	"fmt"
)

type Node interface {
	fmt.Stringer
	Accept(visitor Visitor) any
}

type QueryStmt struct {
	Stmt        Node
	IntoOutFile *string
	Format      Node
}

// Accept implements Node.
func (q *QueryStmt) Accept(visitor Visitor) any {
	panic("unimplemented")
}

// String implements Node.
func (q *QueryStmt) String() string {
	panic("unimplemented")
}

var _ Node = (*QueryStmt)(nil)

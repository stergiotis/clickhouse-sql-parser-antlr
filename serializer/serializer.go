package serializer

import (
	"github.com/antlr4-go/antlr/v4"
)

func Serialize(tree antlr.ParseTree) (_ string) {
	visitor := NewSerializerVisitor()
	result := tree.Accept(visitor)
	return result.(string)
}

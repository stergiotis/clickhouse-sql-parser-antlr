package serializer

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/hungtcs/clickhouse-sql-parser/grammar"
)

// 判断是否是分隔符号如 ",;" 等等
func IsDelimiterToken(tree antlr.Tree) bool {
	if node, ok := tree.(antlr.TerminalNode); ok {
		tokenType := node.GetSymbol().GetTokenType()
		switch tokenType {
		case grammar.ClickHouseLexerSEMICOLON, grammar.ClickHouseParserCOMMA:
			return true
		}
	}
	return false
}

package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/hungtcs/clickhouse-sql-parser/grammar"
)

func antlrParse(input string) (_ grammar.IQueryStmtContext, err error) {
	errorListener := NewErrorListener()

	inputStream := antlr.NewInputStream(input)
	lexer := grammar.NewClickHouseLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewClickHouseParser(tokenStream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	queryStmt := p.QueryStmt()

	if errorListener.HasSyntaxError {
		return nil, errorListener.SyntaxErrors
	}

	return queryStmt, nil
}

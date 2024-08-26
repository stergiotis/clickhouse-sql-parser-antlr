package serializer_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hungtcs/clickhouse-sql-parser/grammar"
	"github.com/hungtcs/clickhouse-sql-parser/parser"
	"github.com/hungtcs/clickhouse-sql-parser/serializer"
	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	assert := assert.New(t)

	input := `select "abc" from t0`
	stmt, err := antlrParse(input)
	assert.NoError(err)

	result := serializer.Serialize(stmt)

	fmt.Printf("result: %v\n", result)
}

func TestSerializeSamples(t *testing.T) {
	entries, err := os.ReadDir("../test_samples")
	if err != nil {
		t.Error(err)
	}
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			assert := assert.New(t)
			data, err := os.ReadFile(path.Join("../test_samples/", entry.Name()))
			if err != nil {
				t.Error(err)
			}
			tree, err := antlrParse(string(data))
			if !assert.Nil(err) {
				fmt.Printf("err.Error(): %v\n", err.Error())
				return
			}

			output := serializer.Serialize(tree)

			fmt.Printf("tree: %v\n", output)
		})
	}
}

func TestChangeAST(t *testing.T) {
	assert := assert.New(t)
	tree, err := antlrParse(`with t0 as (select * from users where create_at > '2020-01-01') select * from t0 limit 10000;`)
	assert.Nil(err)

	output := serializer.Serialize(tree)
	fmt.Printf("output: %v\n", output)
}

func antlrParse(input string) (_ grammar.IQueryStmtContext, err error) {
	errorListener := parser.NewErrorListener()

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

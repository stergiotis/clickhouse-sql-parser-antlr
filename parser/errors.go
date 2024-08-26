package parser

import (
	"bytes"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type SyntaxError struct {
	Line   int
	Column int
	Msg    string
}

// Error implements error.
func (s *SyntaxError) Error() string {
	return fmt.Sprintf("line %d:%d: %s", s.Line, s.Column, s.Msg)
}

func NewSyntaxError(line, col int, msg string) *SyntaxError {
	return &SyntaxError{Line: line, Column: col, Msg: msg}
}

type SyntaxErrors []*SyntaxError

// Error implements error.
func (errs SyntaxErrors) Error() string {
	msg := bytes.NewBuffer(nil)
	for _, err := range errs {
		msg.WriteString(err.Error())
		msg.WriteRune('\n')
	}
	return msg.String()
}

type RuntimeError struct {
	cause       error
	StartLine   int
	EndLine     int
	StartColumn int
	EndColumn   int
}

// Error implements error.
func (r *RuntimeError) Error() string {
	return fmt.Sprintf("line %d:%d-%d:%d: %s", r.StartLine, r.StartColumn, r.EndLine, r.EndColumn, r.cause.Error())
}

func (r *RuntimeError) Unwrap() error {
	return r.cause
}

func NewRuntimeError(ctx antlr.ParserRuleContext, cause error) *RuntimeError {
	s := ctx.GetStart()
	sl := s.GetLine()
	sc := s.GetColumn()

	e := ctx.GetStop()
	el := e.GetLine()
	ec := e.GetColumn()

	return &RuntimeError{
		cause:       cause,
		StartLine:   sl,
		EndLine:     el,
		StartColumn: sc,
		EndColumn:   ec,
	}
}

var (
	_ error = (*SyntaxError)(nil)
	_ error = (SyntaxErrors)(nil)
	_ error = (*RuntimeError)(nil)
)

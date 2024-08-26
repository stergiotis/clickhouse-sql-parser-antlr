package parser

import "github.com/antlr4-go/antlr/v4"

type ErrorListener struct {
	SyntaxErrors   SyntaxErrors
	HasSyntaxError bool
}

// ReportAmbiguity implements antlr.ErrorListener.
func (e *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// panic("unimplemented")
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (e *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// panic("unimplemented")
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (e *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	// panic("unimplemented")
}

// SyntaxError implements antlr.ErrorListener.
func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	el.SyntaxErrors = append(el.SyntaxErrors, NewSyntaxError(line, column, msg))
	el.HasSyntaxError = true
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{}
}

var (
	_ antlr.ErrorListener = (*ErrorListener)(nil)
)

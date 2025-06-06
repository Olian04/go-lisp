package evaluator

import (
	"fmt"

	"github.com/Olian04/go-lisp/lisp/ast"
	"github.com/Olian04/go-lisp/lisp/evaluator/builtins"
	"github.com/Olian04/go-lisp/lisp/evaluator/context"
)

type evaluatorState struct {
	statements []ast.Statement
	context    context.EvaluatorContext
}

func Evaluate(statements []ast.Statement, context context.EvaluatorContext) error {
	state := evaluatorState{
		statements: statements,
		context:    context,
	}

	for _, statement := range state.statements {
		var err error
		state, err = evaluateStatement(state, statement)
		if err != nil {
			return err
		}
	}

	return nil
}

func evaluateStatement(state evaluatorState, statement ast.Statement) (evaluatorState, error) {
	switch statement := statement.(type) {
	case ast.Expression:
		return evaluateExpression(state, statement)
	case ast.Literal:
		return evaluateLiteral(state, statement)
	default:
		return state, fmt.Errorf("unknown statement type: %T", statement)
	}
}

func evaluateExpression(state evaluatorState, expression ast.Expression) (evaluatorState, error) {
	switch expression.Identifier {
	case "print":
		builtins.Print(expression.Arguments, state.context)
		return state, nil
	default:
		return state, fmt.Errorf("unknown expression: %s", expression.Identifier)
	}
}

func evaluateLiteral(state evaluatorState, literal ast.Literal) (evaluatorState, error) {
	return state, nil
}

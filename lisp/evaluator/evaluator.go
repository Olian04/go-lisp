package evaluator

import (
	"fmt"

	"github.com/Olian04/go-lisp/lisp/evaluator/builtins"
	"github.com/Olian04/go-lisp/lisp/evaluator/context"
	"github.com/Olian04/go-lisp/lisp/parser/ast"
)

type evaluatorState struct {
	statements []ast.Statement
	context    context.EvaluatorContext
}

// ast.Program is an array of ast.Statement, which is a union of ast.Expression and ast.Literal
func Evaluate(program ast.Program, context context.EvaluatorContext) error {
	state := evaluatorState{
		statements: program,
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
	switch kind := statement.Kind(); kind {
	case ast.StatementKindExpression:
		return evaluateExpression(state, statement.(ast.Expression))
	case ast.StatementKindLiteral:
		return evaluateLiteral(state, statement.(ast.Literal))
	default:
		return state, fmt.Errorf("unknown statement kind: %s", kind)
	}
}

func evaluateExpression(state evaluatorState, expression ast.Expression) (evaluatorState, error) {
	switch expression.Identifier {
	case "print":
		builtins.Print(expression.Arguments, state.context)
		return state, nil

	case "+":
		fallthrough
	case "add":
		builtins.Add(expression.Arguments, state.context)
		return state, nil

	case "-":
		fallthrough
	case "sub":
		builtins.Sub(expression.Arguments, state.context)
		return state, nil

	case "*":
		fallthrough
	case "mul":
		builtins.Mul(expression.Arguments, state.context)
		return state, nil

	case "/":
		fallthrough
	case "div":
		builtins.Div(expression.Arguments, state.context)
		return state, nil

	case "%":
		fallthrough
	case "mod":
		builtins.Mod(expression.Arguments, state.context)
		return state, nil

	case "**":
		fallthrough
	case "^":
		fallthrough
	case "pow":
		builtins.Pow(expression.Arguments, state.context)
		return state, nil

	case "abs":
		builtins.Abs(expression.Arguments, state.context)
		return state, nil

	case "min":
		builtins.Min(expression.Arguments, state.context)
		return state, nil

	case "max":
		builtins.Max(expression.Arguments, state.context)
		return state, nil

	case "ceil":
		builtins.Ceil(expression.Arguments, state.context)
		return state, nil

	case "floor":
		builtins.Floor(expression.Arguments, state.context)
		return state, nil

	case "round":
		builtins.Round(expression.Arguments, state.context)
		return state, nil

	case "trunc":
		builtins.Trunc(expression.Arguments, state.context)
		return state, nil

	case "sqrt":
		builtins.Sqrt(expression.Arguments, state.context)
		return state, nil

	case "sin":
		builtins.Sin(expression.Arguments, state.context)
		return state, nil

	case "cos":
		builtins.Cos(expression.Arguments, state.context)
		return state, nil

	case "tan":
		builtins.Tan(expression.Arguments, state.context)
		return state, nil

	case "asin":
		builtins.Asin(expression.Arguments, state.context)
		return state, nil

	case "acos":
		builtins.Acos(expression.Arguments, state.context)
		return state, nil

	case "atan":
		builtins.Atan(expression.Arguments, state.context)
		return state, nil

	default:
		return state, fmt.Errorf("unknown expression: %s", expression.Identifier)
	}
}

func evaluateLiteral(state evaluatorState, literal ast.Literal) (evaluatorState, error) {
	return state, nil
}

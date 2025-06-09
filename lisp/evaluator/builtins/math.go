package builtins

import (
	"math"

	"github.com/Olian04/go-lisp/lisp/evaluator/context"
	"github.com/Olian04/go-lisp/lisp/evaluator/util"
	"github.com/Olian04/go-lisp/lisp/parser/ast"
)

func Add(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertAtLeastArgs(args, 1)
	sum := 0.0
	for _, arg := range args {
		sum += arg.(ast.Literal).Value.(float64)
	}
	return sum
}

func Sub(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertAtLeastArgs(args, 1)
	sum := 0.0
	for _, arg := range args {
		sum -= arg.(ast.Literal).Value.(float64)
	}
	return sum
}

func Mul(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertAtLeastArgs(args, 1)
	sum := 0.0
	for _, arg := range args {
		sum *= arg.(ast.Literal).Value.(float64)
	}
	return sum
}

func Div(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertAtLeastArgs(args, 1)
	sum := 0.0
	for _, arg := range args {
		sum /= arg.(ast.Literal).Value.(float64)
	}
	return sum
}

func Mod(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 2)
	return math.Mod(args[0].(ast.Literal).Value.(float64), args[1].(ast.Literal).Value.(float64))
}

func Pow(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 2)
	return math.Pow(args[0].(ast.Literal).Value.(float64), args[1].(ast.Literal).Value.(float64))
}

func Abs(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Abs(args[0].(ast.Literal).Value.(float64))
}

func Min(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertAtLeastArgs(args, 1)
	min := 0.0
	for _, arg := range args {
		min = math.Min(min, arg.(ast.Literal).Value.(float64))
	}
	return min
}

func Max(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertAtLeastArgs(args, 1)
	max := 0.0
	for _, arg := range args {
		max = math.Max(max, arg.(ast.Literal).Value.(float64))
	}
	return max
}

func Ceil(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Ceil(args[0].(ast.Literal).Value.(float64))
}

func Floor(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Floor(args[0].(ast.Literal).Value.(float64))
}

func Round(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Round(args[0].(ast.Literal).Value.(float64))
}

func Trunc(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Trunc(args[0].(ast.Literal).Value.(float64))
}

func Sqrt(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Sqrt(args[0].(ast.Literal).Value.(float64))
}

func Sin(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Sin(args[0].(ast.Literal).Value.(float64))
}

func Cos(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Cos(args[0].(ast.Literal).Value.(float64))
}

func Tan(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Tan(args[0].(ast.Literal).Value.(float64))
}

func Asin(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Asin(args[0].(ast.Literal).Value.(float64))
}

func Acos(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Acos(args[0].(ast.Literal).Value.(float64))
}

func Atan(args []ast.Statement, context context.EvaluatorContext) float64 {
	util.AssertExactArgs(args, 1)
	return math.Atan(args[0].(ast.Literal).Value.(float64))
}

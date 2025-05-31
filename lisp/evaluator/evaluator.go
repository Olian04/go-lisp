package evaluator

import "github.com/Olian04/go-lisp/lisp/ast"

type Evaluator struct {
	program *ast.Program
}

func New(program *ast.Program) *Evaluator {
	return &Evaluator{program: program}
}

func (e *Evaluator) Evaluate() (any, error) {
	return nil, nil
}

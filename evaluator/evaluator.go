package evaluator

import (
	"github.com/DougAlves/interpreter/ast"
	"github.com/DougAlves/interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	}
	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, s := range statements {
		result = Eval(s)
	}
	return result
}

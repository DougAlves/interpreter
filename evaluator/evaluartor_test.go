package evaluator

import (
	"testing"

	"github.com/DougAlves/interpreter/lexer"
	"github.com/DougAlves/interpreter/object"
	"github.com/DougAlves/interpreter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input string
		expected int64
	} {
		{"5", 5},
		{"10", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testInteger(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}


func testInteger(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer, got=%T (%+V)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrog value, expected = %d, but got=%T (%+V)", expected, obj, obj)
	}
	return true
}


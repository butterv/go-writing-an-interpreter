package evaluator

import (
	"github.com/istsh/go-writing-an-interpreter/monkey/ast"
	"github.com/istsh/go-writing-an-interpreter/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}

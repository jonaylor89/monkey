package evaluator

import (
	"MonkeyInterpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len": object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last": object.GetBuiltinByName("last"),
	"rest": object.GetBuiltinByName("rest"),
	"push": object.GetBuiltinByName("push"),
	"puts": object.GetBuiltinByName("puts"),
	"exit": object.GetBuiltinByName("exit"),
}

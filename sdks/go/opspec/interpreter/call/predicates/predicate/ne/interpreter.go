package ne

import (
	"github.com/opctl/opctl/sdks/go/model"
	stringPkg "github.com/opctl/opctl/sdks/go/opspec/interpreter/string"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fakeInterpreter.go --fake-name FakeInterpreter ./ Interpreter

type Interpreter interface {
	Interpret(
		expressions []interface{},
		opHandle model.DataHandle,
		scope map[string]*model.Value,
	) (bool, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter() Interpreter {
	return &_interpreter{
		stringInterpreter: stringPkg.NewInterpreter(),
	}
}

type _interpreter struct {
	stringInterpreter stringPkg.Interpreter
}

func (itp _interpreter) Interpret(
	expressions []interface{},
	opHandle model.DataHandle,
	scope map[string]*model.Value,
) (bool, error) {
	var itemsAsStrings []string
	for _, expression := range expressions {
		// interpret items as strings since everything is coercible to string
		item, err := itp.stringInterpreter.Interpret(scope, expression, opHandle)
		if nil != err {
			return false, err
		}
		currentItemAsString := *item.String

		for _, prevItemAsString := range itemsAsStrings {
			// if any previously seen item equals current item predicate is false.
			if prevItemAsString == currentItemAsString {
				return false, nil
			}
		}

		itemsAsStrings = append(itemsAsStrings, currentItemAsString)
	}
	return true, nil
}

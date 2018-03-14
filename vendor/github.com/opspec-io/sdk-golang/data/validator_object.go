package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/xeipuuv/gojsonschema"
	"strings"
)

// validateObject validates a value against an object parameter
func (vdt _validator) validateObject(
	value *model.Value,
	constraints *model.ObjectConstraints,
) []error {
	valueAsObject, err := vdt.coerce.ToObject(value)
	if nil != err {
		return []error{err}
	}

	// guard no constraints
	if nil != constraints {
		errs := []error{}

		// perform validations supported by gojsonschema
		constraintsJSONBytes, err := json.Marshal(constraints)
		if err != nil {
			// handle syntax errors specially
			return append(
				errs,
				fmt.Errorf("error interpreting constraints; the pkg likely has a syntax error. Details: %v", err.Error()),
			)
		}

		valueJSONBytes, err := json.Marshal(valueAsObject.Object)
		if err != nil {
			// handle syntax errors specially
			return append(
				errs,
				fmt.Errorf("error validating parameter. Details: %v", err.Error()),
			)
		}

		result, err := gojsonschema.Validate(
			gojsonschema.NewBytesLoader(constraintsJSONBytes),
			gojsonschema.NewBytesLoader(valueJSONBytes),
		)
		if err != nil {
			// handle syntax errors specially
			return append(
				errs,
				fmt.Errorf("error validating param. Details: %v", err.Error()),
			)
		}

		for _, errString := range result.Errors() {
			// enum validation errors include `(root) ` prefix we don't want
			errs = append(errs, errors.New(strings.TrimPrefix(errString.Description(), "(root) ")))
		}

		return errs
	}

	return []error{}
}

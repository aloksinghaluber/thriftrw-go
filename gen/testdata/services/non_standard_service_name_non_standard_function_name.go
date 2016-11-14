// Code generated by thriftrw v1.0.0
// @generated

package services

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type NonStandardServiceName_NonStandardFunctionName_Args struct{}

func (v *NonStandardServiceName_NonStandardFunctionName_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *NonStandardServiceName_NonStandardFunctionName_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *NonStandardServiceName_NonStandardFunctionName_Args) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("NonStandardServiceName_NonStandardFunctionName_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *NonStandardServiceName_NonStandardFunctionName_Args) MethodName() string {
	return "non_standard_function_name"
}

func (v *NonStandardServiceName_NonStandardFunctionName_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var NonStandardServiceName_NonStandardFunctionName_Helper = struct {
	Args           func() *NonStandardServiceName_NonStandardFunctionName_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*NonStandardServiceName_NonStandardFunctionName_Result, error)
	UnwrapResponse func(*NonStandardServiceName_NonStandardFunctionName_Result) error
}{}

func init() {
	NonStandardServiceName_NonStandardFunctionName_Helper.Args = func() *NonStandardServiceName_NonStandardFunctionName_Args {
		return &NonStandardServiceName_NonStandardFunctionName_Args{}
	}
	NonStandardServiceName_NonStandardFunctionName_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	NonStandardServiceName_NonStandardFunctionName_Helper.WrapResponse = func(err error) (*NonStandardServiceName_NonStandardFunctionName_Result, error) {
		if err == nil {
			return &NonStandardServiceName_NonStandardFunctionName_Result{}, nil
		}
		return nil, err
	}
	NonStandardServiceName_NonStandardFunctionName_Helper.UnwrapResponse = func(result *NonStandardServiceName_NonStandardFunctionName_Result) (err error) {
		return
	}
}

type NonStandardServiceName_NonStandardFunctionName_Result struct{}

func (v *NonStandardServiceName_NonStandardFunctionName_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *NonStandardServiceName_NonStandardFunctionName_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *NonStandardServiceName_NonStandardFunctionName_Result) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("NonStandardServiceName_NonStandardFunctionName_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *NonStandardServiceName_NonStandardFunctionName_Result) MethodName() string {
	return "non_standard_function_name"
}

func (v *NonStandardServiceName_NonStandardFunctionName_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

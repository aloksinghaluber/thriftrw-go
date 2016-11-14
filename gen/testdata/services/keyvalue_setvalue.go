// Code generated by thriftrw v1.0.0
// @generated

package services

import (
	"fmt"
	"go.uber.org/thriftrw/gen/testdata/unions"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type KeyValue_SetValue_Args struct {
	Key   *Key                   `json:"key,omitempty"`
	Value *unions.ArbitraryValue `json:"value,omitempty"`
}

func (v *KeyValue_SetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Key != nil {
		w, err = v.Key.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Value != nil {
		w, err = v.Value.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *KeyValue_SetValue_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Value, err = _ArbitraryValue_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *KeyValue_SetValue_Args) String() string {
	var fields [2]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	if v.Value != nil {
		fields[i] = fmt.Sprintf("Value: %v", v.Value)
		i++
	}
	return fmt.Sprintf("KeyValue_SetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *KeyValue_SetValue_Args) MethodName() string {
	return "setValue"
}

func (v *KeyValue_SetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var KeyValue_SetValue_Helper = struct {
	Args           func(key *Key, value *unions.ArbitraryValue) *KeyValue_SetValue_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*KeyValue_SetValue_Result, error)
	UnwrapResponse func(*KeyValue_SetValue_Result) error
}{}

func init() {
	KeyValue_SetValue_Helper.Args = func(key *Key, value *unions.ArbitraryValue) *KeyValue_SetValue_Args {
		return &KeyValue_SetValue_Args{Key: key, Value: value}
	}
	KeyValue_SetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	KeyValue_SetValue_Helper.WrapResponse = func(err error) (*KeyValue_SetValue_Result, error) {
		if err == nil {
			return &KeyValue_SetValue_Result{}, nil
		}
		return nil, err
	}
	KeyValue_SetValue_Helper.UnwrapResponse = func(result *KeyValue_SetValue_Result) (err error) {
		return
	}
}

type KeyValue_SetValue_Result struct{}

func (v *KeyValue_SetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *KeyValue_SetValue_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *KeyValue_SetValue_Result) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("KeyValue_SetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *KeyValue_SetValue_Result) MethodName() string {
	return "setValue"
}

func (v *KeyValue_SetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

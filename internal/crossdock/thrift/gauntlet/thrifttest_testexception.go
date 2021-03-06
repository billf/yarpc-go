// Code generated by thriftrw v1.0.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ThriftTest_TestException_Args struct {
	Arg *string `json:"arg,omitempty"`
}

func (v *ThriftTest_TestException_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg != nil {
		w, err = wire.NewValueString(*(v.Arg)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestException_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ThriftTest_TestException_Args) String() string {
	var fields [1]string
	i := 0
	if v.Arg != nil {
		fields[i] = fmt.Sprintf("Arg: %v", *(v.Arg))
		i++
	}
	return fmt.Sprintf("ThriftTest_TestException_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestException_Args) MethodName() string {
	return "testException"
}

func (v *ThriftTest_TestException_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ThriftTest_TestException_Helper = struct {
	Args           func(arg *string) *ThriftTest_TestException_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*ThriftTest_TestException_Result, error)
	UnwrapResponse func(*ThriftTest_TestException_Result) error
}{}

func init() {
	ThriftTest_TestException_Helper.Args = func(arg *string) *ThriftTest_TestException_Args {
		return &ThriftTest_TestException_Args{Arg: arg}
	}
	ThriftTest_TestException_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *Xception:
			return true
		default:
			return false
		}
	}
	ThriftTest_TestException_Helper.WrapResponse = func(err error) (*ThriftTest_TestException_Result, error) {
		if err == nil {
			return &ThriftTest_TestException_Result{}, nil
		}
		switch e := err.(type) {
		case *Xception:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ThriftTest_TestException_Result.Err1")
			}
			return &ThriftTest_TestException_Result{Err1: e}, nil
		}
		return nil, err
	}
	ThriftTest_TestException_Helper.UnwrapResponse = func(result *ThriftTest_TestException_Result) (err error) {
		if result.Err1 != nil {
			err = result.Err1
			return
		}
		return
	}
}

type ThriftTest_TestException_Result struct {
	Err1 *Xception `json:"err1,omitempty"`
}

func (v *ThriftTest_TestException_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Err1 != nil {
		w, err = v.Err1.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestException_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Xception_Read(w wire.Value) (*Xception, error) {
	var v Xception
	err := v.FromWire(w)
	return &v, err
}

func (v *ThriftTest_TestException_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Err1, err = _Xception_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Err1 != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("ThriftTest_TestException_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *ThriftTest_TestException_Result) String() string {
	var fields [1]string
	i := 0
	if v.Err1 != nil {
		fields[i] = fmt.Sprintf("Err1: %v", v.Err1)
		i++
	}
	return fmt.Sprintf("ThriftTest_TestException_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestException_Result) MethodName() string {
	return "testException"
}

func (v *ThriftTest_TestException_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

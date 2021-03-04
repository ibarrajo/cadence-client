// Copyright (c) 2017-2021 Uber Technologies Inc.
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

// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// WorkflowService_GetClusterInfo_Args represents the arguments for the WorkflowService.GetClusterInfo function.
//
// The arguments for GetClusterInfo are sent and received over the wire as this struct.
type WorkflowService_GetClusterInfo_Args struct {
}

// ToWire translates a WorkflowService_GetClusterInfo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_GetClusterInfo_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a WorkflowService_GetClusterInfo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_GetClusterInfo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_GetClusterInfo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_GetClusterInfo_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_GetClusterInfo_Args
// struct.
func (v *WorkflowService_GetClusterInfo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("WorkflowService_GetClusterInfo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_GetClusterInfo_Args match the
// provided WorkflowService_GetClusterInfo_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_GetClusterInfo_Args) Equals(rhs *WorkflowService_GetClusterInfo_Args) bool {

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "GetClusterInfo" for this struct.
func (v *WorkflowService_GetClusterInfo_Args) MethodName() string {
	return "GetClusterInfo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_GetClusterInfo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_GetClusterInfo_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.GetClusterInfo
// function.
var WorkflowService_GetClusterInfo_Helper = struct {
	// Args accepts the parameters of GetClusterInfo in-order and returns
	// the arguments struct for the function.
	Args func() *WorkflowService_GetClusterInfo_Args

	// IsException returns true if the given error can be thrown
	// by GetClusterInfo.
	//
	// An error can be thrown by GetClusterInfo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for GetClusterInfo
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// GetClusterInfo into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by GetClusterInfo
	//
	//   value, err := GetClusterInfo(args)
	//   result, err := WorkflowService_GetClusterInfo_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from GetClusterInfo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.ClusterInfo, error) (*WorkflowService_GetClusterInfo_Result, error)

	// UnwrapResponse takes the result struct for GetClusterInfo
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if GetClusterInfo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WorkflowService_GetClusterInfo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_GetClusterInfo_Result) (*shared.ClusterInfo, error)
}{}

func init() {
	WorkflowService_GetClusterInfo_Helper.Args = func() *WorkflowService_GetClusterInfo_Args {
		return &WorkflowService_GetClusterInfo_Args{}
	}

	WorkflowService_GetClusterInfo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.InternalServiceError:
			return true
		case *shared.ServiceBusyError:
			return true
		default:
			return false
		}
	}

	WorkflowService_GetClusterInfo_Helper.WrapResponse = func(success *shared.ClusterInfo, err error) (*WorkflowService_GetClusterInfo_Result, error) {
		if err == nil {
			return &WorkflowService_GetClusterInfo_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_GetClusterInfo_Result.InternalServiceError")
			}
			return &WorkflowService_GetClusterInfo_Result{InternalServiceError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_GetClusterInfo_Result.ServiceBusyError")
			}
			return &WorkflowService_GetClusterInfo_Result{ServiceBusyError: e}, nil
		}

		return nil, err
	}
	WorkflowService_GetClusterInfo_Helper.UnwrapResponse = func(result *WorkflowService_GetClusterInfo_Result) (success *shared.ClusterInfo, err error) {
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WorkflowService_GetClusterInfo_Result represents the result of a WorkflowService.GetClusterInfo function call.
//
// The result of a GetClusterInfo execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WorkflowService_GetClusterInfo_Result struct {
	// Value returned by GetClusterInfo after a successful execution.
	Success              *shared.ClusterInfo          `json:"success,omitempty"`
	InternalServiceError *shared.InternalServiceError `json:"internalServiceError,omitempty"`
	ServiceBusyError     *shared.ServiceBusyError     `json:"serviceBusyError,omitempty"`
}

// ToWire translates a WorkflowService_GetClusterInfo_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_GetClusterInfo_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_GetClusterInfo_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ClusterInfo_Read(w wire.Value) (*shared.ClusterInfo, error) {
	var v shared.ClusterInfo
	err := v.FromWire(w)
	return &v, err
}

func _InternalServiceError_Read(w wire.Value) (*shared.InternalServiceError, error) {
	var v shared.InternalServiceError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_GetClusterInfo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_GetClusterInfo_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_GetClusterInfo_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_GetClusterInfo_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ClusterInfo_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_GetClusterInfo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_GetClusterInfo_Result
// struct.
func (v *WorkflowService_GetClusterInfo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}

	return fmt.Sprintf("WorkflowService_GetClusterInfo_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_GetClusterInfo_Result match the
// provided WorkflowService_GetClusterInfo_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_GetClusterInfo_Result) Equals(rhs *WorkflowService_GetClusterInfo_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "GetClusterInfo" for this struct.
func (v *WorkflowService_GetClusterInfo_Result) MethodName() string {
	return "GetClusterInfo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_GetClusterInfo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

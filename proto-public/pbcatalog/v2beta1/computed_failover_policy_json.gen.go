// Code generated by protoc-json-shim. DO NOT EDIT.
package catalogv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for ComputedFailoverPolicy
func (this *ComputedFailoverPolicy) MarshalJSON() ([]byte, error) {
	str, err := ComputedFailoverPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ComputedFailoverPolicy
func (this *ComputedFailoverPolicy) UnmarshalJSON(b []byte) error {
	return ComputedFailoverPolicyUnmarshaler.Unmarshal(b, this)
}

var (
	ComputedFailoverPolicyMarshaler   = &protojson.MarshalOptions{}
	ComputedFailoverPolicyUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)

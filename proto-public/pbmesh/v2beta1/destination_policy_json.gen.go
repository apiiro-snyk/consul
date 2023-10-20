// Code generated by protoc-json-shim. DO NOT EDIT.
package meshv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for DestinationPolicy
func (this *DestinationPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationPolicy
func (this *DestinationPolicy) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for DestinationConfig
func (this *DestinationConfig) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationConfig
func (this *DestinationConfig) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LocalityPrioritization
func (this *LocalityPrioritization) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityPrioritization
func (this *LocalityPrioritization) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LoadBalancer
func (this *LoadBalancer) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancer
func (this *LoadBalancer) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for RingHashConfig
func (this *RingHashConfig) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RingHashConfig
func (this *RingHashConfig) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for LeastRequestConfig
func (this *LeastRequestConfig) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LeastRequestConfig
func (this *LeastRequestConfig) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for HashPolicy
func (this *HashPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HashPolicy
func (this *HashPolicy) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

// MarshalJSON is a custom marshaler for CookieConfig
func (this *CookieConfig) MarshalJSON() ([]byte, error) {
	str, err := DestinationPolicyMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CookieConfig
func (this *CookieConfig) UnmarshalJSON(b []byte) error {
	return DestinationPolicyUnmarshaler.Unmarshal(b, this)
}

var (
	DestinationPolicyMarshaler   = &protojson.MarshalOptions{}
	DestinationPolicyUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: private/pbresource/v1/tombstone.proto

package resourcev1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *Tombstone) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *Tombstone) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

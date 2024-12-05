// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1beta1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for TrafficLabel
func (this *TrafficLabel) MarshalJSON() ([]byte, error) {
	str, err := TrafficLabelMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficLabel
func (this *TrafficLabel) UnmarshalJSON(b []byte) error {
	return TrafficLabelUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficLabelRule
func (this *TrafficLabelRule) MarshalJSON() ([]byte, error) {
	str, err := TrafficLabelMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficLabelRule
func (this *TrafficLabelRule) UnmarshalJSON(b []byte) error {
	return TrafficLabelUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LabelVar
func (this *LabelVar) MarshalJSON() ([]byte, error) {
	str, err := TrafficLabelMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LabelVar
func (this *LabelVar) UnmarshalJSON(b []byte) error {
	return TrafficLabelUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPLabelRule
func (this *HTTPLabelRule) MarshalJSON() ([]byte, error) {
	str, err := TrafficLabelMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPLabelRule
func (this *HTTPLabelRule) UnmarshalJSON(b []byte) error {
	return TrafficLabelUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TrafficLabelMarshaler   = &jsonpb.Marshaler{}
	TrafficLabelUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)

package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Int32 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int32 optionalInt32

type optionalInt32 []int32

const (
	valueKeyInt32 = iota
)

// Of wraps the value in an Optional.
func OfInt32(value int32) Int32 {
	return Int32{valueKeyInt32: value}
}

func OfInt32Ptr(ptr *int32) Int32 {
	if ptr == nil {
		return EmptyInt32()
	} else {
		return OfInt32(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt32() Int32 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int32) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int32) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int32) If(f func(value int32)) {
	if o.IsPresent() {
		f(o[valueKeyInt32])
	}
}

func (o Int32) ElseFunc(f func() int32) (value int32) {
	if o.IsPresent() {
		o.If(func(v int32) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int32) Else(elseValue int32) (value int32) {
	return o.ElseFunc(func() int32 { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Int32) ElseZero() (value int32) {
	var zero int32
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int32) String() string {
	if o.IsPresent() {
		var value int32
		o.If(func(v int32) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Int32) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Int32) UnmarshalJSON(data []byte) error {
	var v int32
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfInt32(v)
	return nil
}

func (o Int32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Int32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int32
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt32(v)
	return nil
}

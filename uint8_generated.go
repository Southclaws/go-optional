package optional

import (
	"fmt"
	"time"
)

var _Uint8 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint8 optionalUint8

type optionalUint8 []uint8

const (
	valueKeyUint8 = iota
)

// Of wraps the value in an Optional.
func OfUint8(value uint8) Uint8 {
	return Uint8{valueKeyUint8: value}
}

func OfUint8Ptr(ptr *uint8) Uint8 {
	if ptr == nil {
		return EmptyUint8()
	} else {
		return OfUint8(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUint8() Uint8 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint8) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint8) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint8) If(f func(value uint8)) {
	if o.IsPresent() {
		f(o[valueKeyUint8])
	}
}

func (o Uint8) ElseFunc(f func() uint8) (value uint8) {
	if o.IsPresent() {
		o.If(func(v uint8) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint8) Else(elseValue uint8) (value uint8) {
	return o.ElseFunc(func() uint8 { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint8) String() string {
	if o.IsPresent() {
		var value uint8
		o.If(func(v uint8) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

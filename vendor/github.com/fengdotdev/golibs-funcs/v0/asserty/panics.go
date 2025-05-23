package asserty

import (
	"fmt"
	"reflect"
)

// True checks if the value is true. If not, it panics.
func True(value bool) {
	if !value {
		panic("Assertion failed, expected true but got false")
	}
}

// TrueWithMessage checks if the value is true. If not, it panics with the provided message.
func TrueWithMessage(value bool, message string, args ...any) {

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if !value {
		panic(m)
	}
}

// False checks if the value is false. If not, it panics.
func False(value bool) {
	if value {
		panic("Assertion failed, expected false but got true")
	}
}

// FalseWithMessage checks if the value is false. If not, it panics with the provided message.
func FalseWithMessage(value bool, message string, args ...any) {
	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if value {
		panic(m)
	}
}

// Equal checks if the value is equal to the expected value. If not, it panics.
func Equal(value interface{}, expected interface{}) {
	if !reflect.DeepEqual(value, expected) {
		panic(fmt.Sprintf("Assertion failed, expected %v but got %v", expected, value))
	}
}

// EqualWithMessage checks if the value is equal to the expected value. If not, it panics with the provided message.
func EqualWithMessage(value interface{}, expected interface{}, message string, args ...any) {
	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if !reflect.DeepEqual(value, expected) {
		panic(m)
	}
}

// Nil checks if the value is nil. If not, it panics.
func Nil(value interface{}) {
	if value != nil {
		panic("Assertion failed, expected nil but got non-nil value")
	}
}

// NilWithMessage checks if the value is nil. If not, it panics with the provided message.
func NilWithMessage(value interface{}, message string, args ...any) {
	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if value != nil {
		panic(m)
	}
}

// Err expects an error. If the error is nil, it panics.
func Err(err error) {
	if err == nil {
		panic("Assertion failed, expected error but got nil")
	}
}

// ErrWithMessage expects an error. If the error is nil, it panics with the provided message.
func ErrWithMessage(err error, message string, args ...any) {
	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if err == nil {
		panic(m)
	}
}

// NoError checks if the error is nil. If not, it panics with the error message.
func NoError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Assertion failed, expected no error but got: %v", err))
	}
}

// NoErrorWithMessage checks if the error is nil. If not, it panics with the provided message.
func NoErrorWithMessage(err error, message string, args ...any) {
	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if err != nil {
		panic(m)
	}
}

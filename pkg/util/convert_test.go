package util

import (
	"reflect"
	"testing"
	"time"
)

// Test ToString function.
func TestToString(t *testing.T) {
	mxypltlyx := struct {
		name string
		flag bool
		i    int
		v32  uint32
		v64  uint64
		f    any
		f32  float32
		f64  float64
	}{
		name: "mixie",
		flag: true,
		i:    42,
		v32:  uint32(42),
		v64:  uint64(42),
		f:    3.1415,
		f32:  float32(3.1415),
		f64:  float64(3.14159265),
	}

	var chani chan int

	units := []struct {
		name     string
		value    any
		expected string
	}{
		{
			name:     "string",
			value:    "value-1234",
			expected: "value-1234",
		},
		{
			name:     "bool true",
			value:    true,
			expected: "true",
		},
		{
			name:     "bool false",
			value:    false,
			expected: "false",
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: "7",
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: "0",
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: "255",
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: "2147483647",
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: "9223372036854775807",
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: "0",
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: "255",
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: "0",
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: "65535",
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: "0",
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: "4294967295",
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: "0",
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: "18446744073709551615",
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: "-128",
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: "127",
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: "-32768",
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: "32767",
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: "-2147483648",
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: "2147483647",
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: "-9223372036854775808",
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: "9223372036854775807",
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: "42",
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: "5150",
		},
		{
			name:     "float",
			value:    3.14,
			expected: "3.14",
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: "3.1415",
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: "3.14159265",
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: "65",
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: "[104 105 32]",
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: "90",
		},
		{
			name:     "string array",
			value:    []string{"strings", "vs", "humanity"},
			expected: "[strings vs humanity]",
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: "{data}",
		},
		{
			name:     "mixed struct",
			value:    mxypltlyx,
			expected: "{mixie true 42 42 42 3.1415 3.1415 3.14159265}",
		},
		{
			name:     "duration seconds 0 test",
			value:    time.Duration(0) * time.Second,
			expected: "0s",
		},
		{
			name:     "duration 1 second test",
			value:    time.Duration(1) * time.Second,
			expected: "1s",
		},
		{
			name:     "duration 60 seconds test",
			value:    time.Duration(60) * time.Second,
			expected: "1m0s",
		},
		{
			name:     "duration 300 seconds test",
			value:    time.Duration(300) * time.Second,
			expected: "5m0s",
		},
		{
			name:     "duration 3605 seconds test",
			value:    time.Duration(3605) * time.Second,
			expected: "1h0m5s",
		},
		{
			name:     "2chan test",
			value:    chani,
			expected: "<nil>",
		},
		{
			name: "map test",
			value: map[string]int{"a": 1, "b": 2, "c": 3,
				"z": 26,
			},
			expected: "map[a:1 b:2 c:3 z:26]",
		},
		{
			name:     "func test",
			value:    0xdeadbeef,
			expected: "3735928559",
		},
		{
			name:     "any test",
			value:    any("anyhoo"),
			expected: "anyhoo",
		},
		{
			name:     "interface test",
			value:    interface{}("interface this"),
			expected: "interface this",
		},
	}

	for _, step := range units {
		v := ToString(step.value)
		if v != step.expected {
			t.Errorf("test %v expected %v, got %v", step.name,
				step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.String {
			t.Errorf("test %v expected string, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToString.

// Test ToBoolean function.
func TestToBoolean(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected bool
		errors   bool
	}{
		{
			name:     "bool true",
			value:    true,
			expected: true,
			errors:   false,
		},
		{
			name:     "bool false",
			value:    false,
			expected: false,
			errors:   false,
		},
		{
			name:     "string",
			value:    "value-1234",
			expected: false,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: false,
			errors:   true,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: false,
			errors:   true,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: false,
			errors:   true,
		},
		{
			name:     "float",
			value:    3.1415,
			expected: false,
			errors:   true,
		},
		{
			name:     "float32",
			value:    float32(42.42),
			expected: false,
			errors:   true,
		},
		{
			name:     "float64",
			value:    float64(7.77),
			expected: false,
			errors:   true,
		},
		{
			name:     "byte",
			value:    byte(0),
			expected: false,
			errors:   false,
		},
		{
			name:     "byte",
			value:    byte(1),
			expected: true,
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi yukio"),
			expected: false,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: false,
			errors:   true,
		},
		{
			name:     "string array",
			value:    []string{"number", "9", "backmasking"},
			expected: false,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: false,
			errors:   true,
		},
		{
			name:     "duration 60 seconds",
			value:    time.Duration(60) * time.Second,
			expected: false,
			errors:   true,
		},
		{
			name:     "duration 8 hours test",
			value:    time.Duration(8) * time.Hour,
			expected: false,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 0),
			expected: false,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]string{"abc": "123"},
			expected: false,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: false,
			errors:   true,
		},
		{
			name:     "any test",
			value:    any("true"),
			expected: true,
			errors:   false,
		},
		{
			name:     "any false test",
			value:    any("false"),
			expected: false,
			errors:   false,
		},
		{
			name:     "any other test",
			value:    any("other"),
			expected: false,
			errors:   true,
		},
		{
			name:     "interface test",
			value:    interface{}("true"),
			expected: true,
			errors:   false,
		},
		{
			name:     "interface false test",
			value:    interface{}("false"),
			expected: false,
			errors:   false,
		},
		{
			name:     "interface other test",
			value:    interface{}("something wicked this way comes"),
			expected: false,
			errors:   true,
		},
	}

	for _, step := range units {
		v, err := ToBoolean(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Bool {
			t.Errorf("test %v expected boolean, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToBoolean.

// Test ToInteger function.
func TestToInteger(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected int
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: 42,
			errors:   false,
		},
		{
			name:     "string",
			value:    "1234",
			expected: 1234,
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: -1,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: -1,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: 7,
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: 0,
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: 255,
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: 2147483647,
			errors:   false,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: 9223372036854775807,
			errors:   false,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: 0,
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: 255,
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: 0,
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: 65535,
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: 0,
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: 4294967295,
			errors:   false,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: 0,
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: -1,
			errors:   true,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: -128,
			errors:   false,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: 127,
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: -32768,
			errors:   false,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: 32767,
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: -2147483648,
			errors:   false,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: 2147483647,
			errors:   false,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: -9223372036854775808,
			errors:   false,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: 9223372036854775807,
			errors:   false,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: 42,
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: 5150,
			errors:   false,
		},
		{
			name:     "float",
			value:    3.14,
			expected: -1,
			errors:   true,
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: -1,
			errors:   true,
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: -1,
			errors:   true,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: 65,
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: -1,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: 90,
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: -1,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: -1,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: -1,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 0),
			expected: -1,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]int{"abcd": 1234},
			expected: -1,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: -1,
			errors:   true,
		},
		{
			name:     "any test",
			value:    any("true"),
			expected: -1,
			errors:   true,
		},
		{
			name:     "any false test",
			value:    any("false"),
			expected: -1,
			errors:   true,
		},
		{
			name:     "any number test",
			value:    any("1234"),
			expected: 1234,
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("intermission"),
			expected: -1,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("false"),
			expected: -1,
			errors:   true,
		},
		{
			name:     "interface valid test",
			value:    interface{}("42"),
			expected: 42,
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToInteger(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		if kind := reflect.ValueOf(v).Kind(); kind != reflect.Int {
			t.Errorf("test %v expected integer, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToInteger.

// Test ToUnsignedInt16 function.
func TestToUnsignedInt16(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected uint16
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: uint16(42),
			errors:   false,
		},
		{
			name:     "string",
			value:    "1234",
			expected: uint16(1234),
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: 0,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: 0,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: uint16(7),
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: uint16(0),
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: uint16(255),
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: 0,
			errors:   true,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: uint16(0),
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: uint16(255),
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: uint16(0),
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: uint16(65535),
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: uint16(0),
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: 0,
			errors:   true,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: uint16(0),
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: uint16(127),
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: uint16(32767),
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: 0,
			errors:   true,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: uint16(42),
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: uint16(5150),
			errors:   false,
		},
		{
			name:     "float",
			value:    3.14,
			expected: 0,
			errors:   true,
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: 0,
			errors:   true,
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: 0,
			errors:   true,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: uint16(65),
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: 0,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: uint16(90),
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: 0,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 0),
			expected: 0,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]int{"answer": 42},
			expected: 0,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: 0,
			errors:   true,
		},
		{
			name:     "any true test",
			value:    any("true"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "any number test",
			value:    any("42"),
			expected: 42,
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("blah"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("false"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface valid test",
			value:    interface{}("7"),
			expected: 7,
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToUnsignedInt16(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Uint16 {
			t.Errorf("%v expected uint16, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToUnsignedInt16.

// Test ToUnsignedInt32 function.
func TestToUnsignedInt32(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected uint32
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: uint32(42),
			errors:   false,
		},
		{
			name:     "string",
			value:    "1234",
			expected: uint32(1234),
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: 0,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: 0,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: uint32(7),
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: uint32(0),
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: uint32(255),
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: uint32(2147483647),
			errors:   false,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: 0,
			errors:   true,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: uint32(0),
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: uint32(255),
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: uint32(0),
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: uint32(65535),
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: uint32(0),
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: uint32(4294967295),
			errors:   false,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: uint32(0),
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: uint32(127),
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: uint32(32767),
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: uint32(2147483647),
			errors:   false,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: 0,
			errors:   true,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: uint32(42),
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: uint32(5150),
			errors:   false,
		},
		{
			name:     "float",
			value:    3.14,
			expected: 0,
			errors:   true,
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: 0,
			errors:   true,
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: 0,
			errors:   true,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: uint32(65),
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: 0,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: uint32(90),
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: 0,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 0),
			expected: 0,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]int{"answer": 42},
			expected: 0,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: 0,
			errors:   true,
		},
		{
			name:     "any true test",
			value:    any("true"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "any number test",
			value:    any("999"),
			expected: 999,
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("error"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("true"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface valid test",
			value:    interface{}("1024"),
			expected: 1024,
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToUnsignedInt32(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Uint32 {
			t.Errorf("test %v expected uint32, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToUnsignedInt32.

// Test ToUnsignedInt64 function.
func TestToUnsignedInt64(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected uint64
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: uint64(42),
			errors:   false,
		},
		{
			name:     "string",
			value:    "1234",
			expected: uint64(1234),
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: 0,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: 0,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: uint64(7),
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: uint64(0),
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: uint64(255),
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: uint64(2147483647),
			errors:   false,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: uint64(9223372036854775807),
			errors:   false,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: uint64(0),
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: uint64(255),
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: uint64(0),
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: uint64(65535),
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: uint64(0),
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: uint64(4294967295),
			errors:   false,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: uint64(0),
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: uint64(18446744073709551615),
			errors:   false,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: uint64(127),
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: uint64(32767),
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: uint64(2147483647),
			errors:   false,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: 0,
			errors:   true,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: uint64(9223372036854775807),
			errors:   false,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: uint64(42),
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: uint64(5150),
			errors:   false,
		},
		{
			name:     "float",
			value:    3.14,
			expected: 0,
			errors:   true,
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: 0,
			errors:   true,
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: 0,
			errors:   true,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: uint64(65),
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: 0,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: uint64(90),
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: 0,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 0),
			expected: 0,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]int{"error": 404},
			expected: 0,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: 0,
			errors:   true,
		},
		{
			name:     "any false test",
			value:    any("false"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "any number test",
			value:    any("2048"),
			expected: 2048,
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("errorneous"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("true"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface valid test",
			value:    interface{}("4096"),
			expected: 4096,
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToUnsignedInt64(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Uint64 {
			t.Errorf("test %v expected uint64, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToUnsignedInt64.

// Test ToFloat32 function.
func TestToFloat32(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected float32
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: float32(42),
			errors:   false,
		},
		{
			name:     "string",
			value:    "1234",
			expected: float32(1234),
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: 0,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: 0,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: float32(7),
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: float32(0),
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: float32(255),
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: float32(2147483647),
			errors:   false,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: float32(9223372036854775807),
			errors:   false,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: float32(0),
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: float32(255),
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: float32(0),
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: float32(65535),
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: float32(0),
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: float32(4294967295),
			errors:   false,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: float32(0),
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: float32(18446744073709551615),
			errors:   false,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: float32(-128),
			errors:   false,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: float32(127),
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: float32(-32768),
			errors:   false,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: float32(32767),
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: float32(-2147483648),
			errors:   false,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: float32(2147483647),
			errors:   false,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: float32(-9223372036854775808),
			errors:   false,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: float32(9223372036854775807),
			errors:   false,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: float32(42),
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: float32(5150),
			errors:   false,
		},
		{
			name:     "float",
			value:    3.14,
			expected: float32(3.14),
			errors:   false,
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: float32(3.1415),
			errors:   false,
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: float32(3.14159265),
			errors:   false,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: float32(65),
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: 0,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: float32(90),
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: 0,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 2),
			expected: 0,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]float32{"pi": float32(3.14)},
			expected: 0,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: 0,
			errors:   true,
		},
		{
			name:     "any true test",
			value:    any("true"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "any number test",
			value:    any("3.1415"),
			expected: float32(3.1415),
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("nope"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("false"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface valid test",
			value:    interface{}("3.1415"),
			expected: float32(3.1415),
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToFloat32(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Float32 {
			t.Errorf("test %v expected float32, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToFloat32.

// Test ToFloat64 function.
func TestToFloat64(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected float64
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: float64(42),
			errors:   false,
		},
		{
			name:     "string",
			value:    "1234",
			expected: float64(1234),
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: 0,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: 0,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: float64(7),
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: float64(0),
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: float64(255),
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: float64(2147483647),
			errors:   false,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: float64(9223372036854775807),
			errors:   false,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: float64(0),
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: float64(255),
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: float64(0),
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: float64(65535),
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: float64(0),
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: float64(4294967295),
			errors:   false,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: float64(0),
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: float64(18446744073709551615),
			errors:   false,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: float64(-128),
			errors:   false,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: float64(127),
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: float64(-32768),
			errors:   false,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: float64(32767),
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: float64(-2147483648),
			errors:   false,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: float64(2147483647),
			errors:   false,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: float64(-9223372036854775808),
			errors:   false,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: float64(9223372036854775807),
			errors:   false,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: float64(42),
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: float64(5150),
			errors:   false,
		},
		{
			name:     "float",
			value:    3.14,
			expected: float64(3.14),
			errors:   false,
		},
		{
			name:     "float32",
			value:    float32(3.1415),
			expected: float64(3.1415),
			errors:   false,
		},
		{
			name:     "float64",
			value:    float64(3.14159265),
			expected: float64(3.14159265),
			errors:   false,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: float64(65),
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: 0,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: float64(90),
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: 0,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: 0,
			errors:   true,
		},
		{
			name:     "chan test",
			value:    make(chan int, 2),
			expected: 0,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]float64{"pi": float64(3.14)},
			expected: 0,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: 0,
			errors:   true,
		},
		{
			name:     "any true test",
			value:    any("false"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "any number test",
			value:    any("3.1415"),
			expected: float64(3.1415),
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("nope"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("true"),
			expected: 0,
			errors:   true,
		},
		{
			name:     "interface valid test",
			value:    interface{}("3.1415"),
			expected: float64(3.1415),
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToFloat64(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v",
				step.name, step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Float64 {
			t.Errorf("test %v expected float64, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToFloat64.

// Test ToTimeDuration function.
func TestToTimeDuration(t *testing.T) {
	units := []struct {
		name     string
		value    any
		expected time.Duration
		errors   bool
	}{
		{
			name:     "int",
			value:    42,
			expected: time.Duration(42) * time.Second,
			errors:   false,
		},
		{
			name:     "string",
			value:    "60",
			expected: time.Duration(60) * time.Second,
			errors:   false,
		},
		{
			name:     "string",
			value:    "-30",
			expected: time.Duration(-30) * time.Second,
			errors:   false,
		},
		{
			name:     "string",
			value:    "90s",
			expected: time.Duration(90) * time.Second,
			errors:   false,
		},
		{
			name:     "string",
			value:    "-120",
			expected: time.Duration(-120) * time.Second,
			errors:   false,
		},
		{
			name:     "string",
			value:    "1h",
			expected: time.Duration(1) * time.Hour,
			errors:   false,
		},
		{
			name:     "bool true",
			value:    true,
			expected: time.Duration(0) * time.Hour,
			errors:   true,
		},
		{
			name:     "bool false",
			value:    false,
			expected: time.Duration(0) * time.Hour,
			errors:   true,
		},
		{
			name:     "int seven",
			value:    int(7),
			expected: time.Duration(7) * time.Second,
			errors:   false,
		},
		{
			name:     "int 0",
			value:    int(0),
			expected: time.Duration(0) * time.Second,
			errors:   false,
		},
		{
			name:     "int 255",
			value:    int(255),
			expected: time.Duration(255) * time.Second,
			errors:   false,
		},
		{
			name:     "int max 32-bit",
			value:    int32(2147483647),
			expected: time.Duration(2147483647) * time.Second,
			errors:   false,
		},
		{
			name:     "int max 64-bit",
			value:    int64(9223372036854775807),
			expected: time.Duration(-1) * time.Second,
			errors:   false,
		},
		{
			name:     "uint8 min",
			value:    uint8(0),
			expected: time.Duration(0) * time.Second,
			errors:   false,
		},
		{
			name:     "uint8 max",
			value:    uint8(255),
			expected: time.Duration(255) * time.Second,
			errors:   false,
		},
		{
			name:     "uint16 min",
			value:    uint16(0),
			expected: time.Duration(0) * time.Second,
			errors:   false,
		},
		{
			name:     "uint16 max",
			value:    uint16(65535),
			expected: time.Duration(65535) * time.Second,
			errors:   false,
		},
		{
			name:     "uint32 min",
			value:    uint32(0),
			expected: time.Duration(0) * time.Second,
			errors:   false,
		},
		{
			name:     "uint32 max",
			value:    uint32(4294967295),
			expected: time.Duration(4294967295) * time.Second,
			errors:   false,
		},
		{
			name:     "uint64 min",
			value:    uint64(0),
			expected: time.Duration(0) * time.Second,
			errors:   false,
		},
		{
			name:     "uint64 max",
			value:    uint64(18446744073709551615),
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "int8 min",
			value:    int8(-128),
			expected: time.Duration(-128) * time.Second,
			errors:   false,
		},
		{
			name:     "int8 max",
			value:    int8(127),
			expected: time.Duration(127) * time.Second,
			errors:   false,
		},
		{
			name:     "int16 min",
			value:    int16(-32768),
			expected: time.Duration(-32768) * time.Second,
			errors:   false,
		},
		{
			name:     "int16 max",
			value:    int16(32767),
			expected: time.Duration(32767) * time.Second,
			errors:   false,
		},
		{
			name:     "int32 min",
			value:    int32(-2147483648),
			expected: time.Duration(-2147483648) * time.Second,
			errors:   false,
		},
		{
			name:     "int32 max",
			value:    int32(2147483647),
			expected: time.Duration(2147483647) * time.Second,
			errors:   false,
		},
		{
			name:     "int64 min",
			value:    int64(-9223372036854775808),
			expected: time.Duration(0) * time.Second,
			errors:   false,
		},
		{
			name:     "int64 max",
			value:    int64(9223372036854775807),
			expected: time.Duration(-1) * time.Second,
			errors:   false,
		},
		{
			name:     "uint32 forty-two",
			value:    uint32(42),
			expected: time.Duration(42) * time.Second,
			errors:   false,
		},
		{
			name:     "uint64 vh",
			value:    uint64(5150),
			expected: time.Duration(5150) * time.Second,
			errors:   false,
		},
		{
			name:     "float",
			value:    7.0,
			expected: time.Duration(7) * time.Second,
			errors:   false,
		},
		{
			name:     "float32",
			value:    float32(3),
			expected: time.Duration(3) * time.Second,
			errors:   false,
		},
		{
			name:     "float64",
			value:    float64(42.0),
			expected: time.Duration(42) * time.Second,
			errors:   false,
		},
		{
			name:     "byte",
			value:    byte(65),
			expected: time.Duration(65) * time.Second,
			errors:   false,
		},
		{
			name:     "byte array",
			value:    []byte("hi "),
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "rune",
			value:    rune(90),
			expected: time.Duration(90) * time.Second,
			errors:   false,
		},
		{
			name:     "string array",
			value:    []string{"just", "some", "words"},
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "struct",
			value:    struct{ n string }{n: "data"},
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "duration 30 seconds test",
			value:    time.Duration(30) * time.Second,
			expected: time.Duration(30) * time.Second,
			errors:   false,
		},
		{
			name:     "duration 1 hour test",
			value:    time.Duration(3600) * time.Second,
			expected: time.Duration(1) * time.Hour,
			errors:   false,
		},
		{
			name:     "chan test",
			value:    make(chan int, 2),
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "map test",
			value:    map[string]float32{"pi": float32(3.14)},
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "func test",
			value:    func() {},
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "any true test",
			value:    any("true"),
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "any floating number test",
			value:    any("60.12"),
			expected: time.Duration(6012) * time.Second / 100,
			errors:   false,
		},
		{
			name:     "any number test",
			value:    any("600"),
			expected: time.Duration(600) * time.Second,
			errors:   false,
		},
		{
			name:     "interface test",
			value:    interface{}("nope"),
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "interface boolean test",
			value:    interface{}("false"),
			expected: time.Duration(0) * time.Second,
			errors:   true,
		},
		{
			name:     "interface valid number test",
			value:    interface{}("3600"),
			expected: time.Duration(1) * time.Hour,
			errors:   false,
		},
		{
			name:     "interface valid floating point test",
			value:    interface{}("3.1415"),
			expected: 31415 * time.Second / 10000,
			errors:   false,
		},
	}

	for _, step := range units {
		v, err := ToTimeDuration(step.value)

		if step.errors {
			if err == nil {
				t.Errorf("test %v expected an error",
					step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("test %v expected no error, got %v",
				step.name, err)
		}

		if v != step.expected {
			t.Errorf("test %v expected %v, got %v", step.name,
				step.expected, v)
		}

		kind := reflect.ValueOf(v).Kind()
		if kind != reflect.Int64 {
			t.Errorf("test %v expected integer, got %v",
				step.name, kind)
		}
	}

} //  End of function  TestToTimeDuration.

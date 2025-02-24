package util

import (
	"fmt"
	"strconv"
	"time"
)

// Returns value as a string.
func ToString(value any) string {
	if s, ok := value.(string); ok {
		return s
	}

	return fmt.Sprintf("%v", value)

} // End of function  ToString.

// Returns value as a boolean.
func ToBoolean(value any) (bool, error) {
	if b, ok := value.(bool); ok {
		return b, nil
	}

	return strconv.ParseBool(fmt.Sprintf("%v", value))

} // End of function  ToBoolean.

// Returns value as an integer.
func ToInteger(value any) (int, error) {
	if v, ok := value.(int); ok {
		return v, nil
	}

	return strconv.Atoi(fmt.Sprintf("%v", value))

} // End of function  ToInteger.

// Returns value as a 16-bit unsigned integer.
func ToUnsignedInt16(value any) (uint16, error) {
	if v, ok := value.(uint16); ok {
		return v, nil
	}

	v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 16)
	return uint16(v), err

} // End of function  ToUnsignedInt16.

// Returns value as a 32-bit unsigned integer.
func ToUnsignedInt32(value any) (uint32, error) {
	if v, ok := value.(uint32); ok {
		return v, nil
	}

	v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 32)
	return uint32(v), err

} // End of function  ToUnsignedInt32.

// Returns value as a 64-bit unsigned integer.
func ToUnsignedInt64(value any) (uint64, error) {
	if v, ok := value.(uint64); ok {
		return v, nil
	}

	v, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
	return uint64(v), err

} // End of function  ToUnsignedInt64.

// Returns value as a 32-bit floating point number.
func ToFloat32(value any) (float32, error) {
	if v, ok := value.(float32); ok {
		return v, nil
	}

	v, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 32)
	return float32(v), err

} // End of function  ToFloat32.

// Returns value as a 64-bit floating point number.
func ToFloat64(value any) (float64, error) {
	if v, ok := value.(float64); ok {
		return v, nil
	}

	v, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	return float64(v), err

} // End of function  ToFloat64.

// Returns value as a time duration.
func ToTimeDuration(value any) (time.Duration, error) {
	if dur, ok := value.(time.Duration); ok {
		return dur, nil
	}

	valueStr := fmt.Sprintf("%v", value)
	if v, err := strconv.Atoi(valueStr); err == nil {
		// `value` string is an integer and we default to
		// using integer value as seconds.
		return time.Duration(v) * time.Second, nil
	}

	if d, err := time.ParseDuration(valueStr); err == nil {
		return d, nil
	}

	return time.ParseDuration(valueStr + "s")

} // End of function  ToTimeDuration.

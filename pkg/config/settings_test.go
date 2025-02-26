package config

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

// Check and compare two maps and return errors if they don't match.
func checkMaps(expected, actual map[string]string) error {
	if reflect.DeepEqual(expected, actual) {
		return nil
	}

	for k, v := range expected {
		if value, ok := actual[k]; !ok {
			return fmt.Errorf("missing setting %v", k)
		} else if v != value {
			return fmt.Errorf("expected %v to be %v, got %v",
				k, v, value)
		}

	}

	// And check that keys in the actual map exist in the expected one.
	for k := range actual {
		if _, ok := expected[k]; !ok {
			return fmt.Errorf("extraneous setting %v", k)
		}
	}

	return fmt.Errorf("expected %v to match %v", actual, expected)

} //  End of function  checkMaps.

// Test LoadSettings function.
func TestLoadSettings(t *testing.T) {
	expectedSettings := map[string]string{
		"TELEX_NAME":                "WRU",
		"TELEX_VERSION":             "0.0.7",
		"TELEX_FLOAT_1":             "0.07",
		"TELEX_FLOAT_2":             "42.42",
		"TELEX_INT_1":               "7",
		"TELEX_INT_2":               "42",
		"TELEX_RETRY_QUEUE_SIZE":    "280",
		"TELEX_BOOL_1":              "true",
		"TELEX_BOOL_2":              "false",
		"TELEX_VALIDATE_TLS_CONFIG": "false",
		"TELEX_SERVICE_ADDRESS":     "127.1.2.3",
		"TELEX_SERVICE_PORT":        "9876",
		"TELEX_INLINE_INT":          "1024",
		"TELEX_FILE":                "/tmp/cable",
		"TELEX_TIMEOUT":             "300",
		"EXT_ID":                    "extension",
		"EXT_PI":                    "3.14",
		"EXT_ANSWER":                "42",
		"EXT_FLAG":                  "false",
		"EXT_COMPLEX":               "a=1,b=two,c=3.14",
		"EXT_COMMENT_1":             "# add some details here #plus this one",
		"EXT_COMMENT_2":             "'# add some details here #plus this one'",
		"EXT_COMMENT_3":             "# add some ####details here",
		"FOO_BAR_BAZ":               "metasyntactic",
	}

	deviceSettings := map[string]string{
		"GRPC_TELEGRAPH_NAME":                   "dev-em",
		"GRPC_TELEGRAPH_VERSION":                "0.4.2",
		"GRPC_TELEGRAPH_DEBUG":                  "true",
		"GRPC_TELEGRAPH_TOKEN":                  "let me inside",
		"GRPC_TELEGRAPH_SERVICE_ADDRESS":        "127.0.0.1",
		"GRPC_TELEGRAPH_SERVICE_PORT":           "9340",
		"GRPC_TELEGRAPH_RETRY_QUEUE_SIZE":       "2048",
		"GRPC_TELEGRAPH_CERT":                   "test/tls/device/telegraph-cert.pem",
		"GRPC_TELEGRAPH_KEY":                    "test/tls/device/telegraph-key.pem",
		"GRPC_TELEGRAPH_SERVICE_CACERT":         "test/tls/service/cacert.pem",
		"GRPC_TELEGRAPH_CONNECT_TIMEOUT":        "30",
		"GRPC_TELEGRAPH_SEND_TIMEOUT":           "60",
		"GRPC_TELEGRAPH_KEEP_ALIVE_TIMEOUT":     "300",
		"GRPC_TELEGRAPH_MAX_SUBSCRIPTION_DELAY": "300",

		// namespaced extensions
		"GRPC_TELEGRAPH_ID":         "extensions",
		"GRPC_TELEGRAPH_EXT_NAME":   "kelex",
		"GRPC_TELEGRAPH_FLOAT_PI":   "3.14",
		"GRPC_TELEGRAPH_INT_ANSWER": "42",
		"GRPC_TELEGRAPH_EXT_ENABLE": "false",

		// non-namespaced extensions
		"ZID":            "non-namespaced squad k",
		"PI":             "3.141592653589793",
		"INT_ANSWER":     "42",
		"REFLAG":         "true",
		"COMPLEX_VALUE":  "a=1,b=two,c=3.14",
		"COMMENT_TEST_1": "# add some details here #plus this one",
		"COMMENT_TEST_2": "'# add some details here #plus this one'",
		"COMMENT_TEST_3": "# add some ####details here",
		"FOO_BAR_BAZ":    "metasyntactic",
	}

	serviceSettings := map[string]string{
		"GRPC_TELEGRAPH_NAME":                      "me:dev-ac",
		"GRPC_TELEGRAPH_VERSION":                   "0.4.2",
		"GRPC_TELEGRAPH_DEBUG":                     "true",
		"GRPC_TELEGRAPH_BIND_ADDRESS":              "127.0.0.1",
		"GRPC_TELEGRAPH_BIND_PORT":                 "9340",
		"GRPC_TELEGRAPH_SERVICE_TYPE":              "station",
		"GRPC_TELEGRAPH_SERVICE_ADDRESS":           "upstream.service.telegraph.local",
		"GRPC_TELEGRAPH_SERVICE_PORT":              "9340",
		"GRPC_TELEGRAPH_CERT":                      "test/tls/service/bundle/service.pem",
		"GRPC_TELEGRAPH_KEY":                       "test/tls/service/bundle/service.pem",
		"GRPC_TELEGRAPH_BOOTSTRAP_CACERTS_PATTERN": "test/tls/bootstrap/*-cacert.pem",
		"GRPC_TELEGRAPH_DEVICE_CACERTS_PATTERN":    "test/tls/device/*-cacert.pem",
		"GRPC_TELEGRAPH_ENABLE_SUBSCRIPTIONS":      "true",
		"GRPC_TELEGRAPH_BUFFER_SIZE":               "4194304",
		"GRPC_TELEGRAPH_MAX_MESSAGE_SIZE":          "4194304",
		"GRPC_TELEGRAPH_MAX_STREAMS":               "255",
		"GRPC_TELEGRAPH_NUM_STREAM_WORKERS":        "100",
		"GRPC_TELEGRAPH_SEND_TIMEOUT":              "60",
		"GRPC_TELEGRAPH_KEEP_ALIVE_TIMEOUT":        "300",

		// namespaced extensions
		"GRPC_TELEGRAPH_ID":         "extensions",
		"GRPC_TELEGRAPH_EXT_NAME":   "saas:y",
		"GRPC_TELEGRAPH_FLOAT_PI":   "3.14",
		"GRPC_TELEGRAPH_INT_ANSWER": "42",
		"GRPC_TELEGRAPH_EXT_ENABLE": "false",

		// non-namespaced extensions
		"ZID":            "non-namespaced TLA",
		"PI":             "3.141592653589793",
		"INT_ANSWER":     "42",
		"REFLAG":         "true",
		"COMPLEX_VALUE":  "a=one,b=2,c=3.14",
		"COMMENT_TEST_1": "# add some details here #plus this one",
		"COMMENT_TEST_2": "'# add some details here #plus this one'",
		"COMMENT_TEST_3": "# add some ####details here",
		"FOO_BAR_BAZ":    "metasyntactic",
	}

	units := []struct {
		name     string
		path     string
		expected map[string]string
		errors   bool
	}{
		{
			name:     "missing file",
			path:     "/tmp/path/to/missing/file/404.env",
			expected: make(map[string]string),
			errors:   true,
		},
		{
			name:     "invalid env file",
			path:     filepath.Join("fixtures", "invalid.env"),
			expected: make(map[string]string),
			errors:   true,
		},
		{
			name:     "empty env file",
			path:     filepath.Join("fixtures", "empty.env"),
			expected: make(map[string]string),
			errors:   false,
		},
		{
			name:     "test env file",
			path:     filepath.Join("fixtures", "test.env"),
			expected: expectedSettings,
			errors:   false,
		},
		{
			name:     "device env file",
			path:     filepath.Join("../..", "config", "device.env"),
			expected: deviceSettings,
			errors:   false,
		},
		{
			name:     "service env file",
			path:     filepath.Join("../..", "config", "service.env"),
			expected: serviceSettings,
			errors:   false,
		},
	}

	for _, step := range units {
		settings, err := LoadSettings(step.path)
		if step.errors {
			if err == nil {
				t.Errorf("test '%v': expected error, got none", step.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("test '%v': unexpected error: %v", step.name, err)
		}

		if err := checkMaps(step.expected, settings); err != nil {
			t.Errorf("test '%v' %v", step.name, err)
		}
	}

} //  End of function  TestLoadSettings.

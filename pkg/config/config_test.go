package config

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

// Config map for various grouped settings.
type SettingsMap map[string]any

// Config test case.
type configTestCase struct {
	name        string
	namespace   string
	path        string
	defaults    map[string]SettingsMap
	expected    map[string]SettingsMap
	errors      bool
	nextensions int
}

// Check an object against a map containing key-value pairs.
func checkObject(obj any, m map[string]any) error {
	value := reflect.ValueOf(obj)
	kind := value.Type()

	for idx := 0; idx < value.NumField(); idx++ {
		k := kind.Field(idx).Name
		expectation, ok := m[k]
		if !ok {
			fmt.Printf("missing field %v in check object\n", k)
			continue
		}

		v := value.Field(idx).Interface()
		if !reflect.DeepEqual(v, expectation) {
			return fmt.Errorf("%v setting expected %v, got %v",
				k, v, expectation)
		}
	}

	return nil

} // End of function  checkObject.

// Returns the default settings in a map of key-value pairs.
func defaultSettings() SettingsMap {
	return map[string]any{
		"Name":        DEFAULT_NAME,
		"Debug":       DEFAULT_DEBUG_FLAG,
		"Version":     DEFAULT_VERSION,
		"ValidateTLS": DEFAULT_VALIDATE_TLS_CONFIG,
		"Cert":        DEFAULT_CERTIFICATE,
		"Key":         DEFAULT_PRIVATE_KEY,
	}

} // End of function  defaultSettings.

// Test makeDefaultSettings function.
func TestMakeDefaultSettings(t *testing.T) {
	expectation := defaultSettings()

	cfg := makeDefaultSettings()
	if err := checkObject(cfg, expectation); err != nil {
		t.Error(err)
	}

} // End of function  TestMakeDefaultSettings.

// Returns the default timeout settings in a map of key-value pairs.
func timeoutSettings() SettingsMap {
	return map[string]any{
		"Connect":              DEFAULT_CONNECT_TIMEOUT,
		"Send":                 DEFAULT_SEND_TIMEOUT,
		"KeepAlive":            DEFAULT_KEEP_ALIVE_TIMEOUT,
		"MaxSubscriptionDelay": DEFAULT_MAX_SUBSCRIPTION_DELAY,
	}

} // End of function  timeoutSettings.

// Test makeDefaultTimeoutSettings function.
func TestMakeDefaultTimeoutSettings(t *testing.T) {
	expectation := timeoutSettings()

	cfg := makeDefaultTimeoutSettings()
	if err := checkObject(cfg, expectation); err != nil {
		t.Error(err)
	}

} // End of function  TestMakeDefaultTimeoutSettings.

// Returns the default device settings in a map of key-value pairs.
func deviceSettings() SettingsMap {
	return map[string]any{
		"Token":          "",
		"ServiceAddress": DEFAULT_SERVICE_ADDRESS,
		"ServicePort":    DEFAULT_SERVICE_PORT_NUMBER,
		"ServiceCACert":  "",
		"RetryQueueSize": DEFAULT_RETRY_QUEUE_SIZE,
	}

} // End of function  deviceSettings.

// Test makeDefaultDeviceSettings function.
func TestMakeDefaultDeviceSettings(t *testing.T) {
	expectation := deviceSettings()

	cfg := makeDefaultDeviceSettings()
	if err := checkObject(cfg, expectation); err != nil {
		t.Error(err)
	}

} // End of function  TestMakeDefaultDeviceSettings.

// Returns the default CA cerificates pattern in a map of key-value pairs.
func caCertificatesPattern() SettingsMap {
	return map[string]any{
		"Bootstrap": DEFAULT_BOOTSTRAP_CACERTS_PATTERN,
		"Device":    DEFAULT_DEVICE_CACERTS_PATTERN,
	}

} // End of function  caCertificatesPattern.

// Test makeDefaultCACertificatesPattern function.
func TestMakeDefaultCACertifiatesPattern(t *testing.T) {
	expectation := caCertificatesPattern()

	cfg := makeDefaultCACertificatesPattern()

	if err := checkObject(cfg, expectation); err != nil {
		t.Error(err)
	}

} // End of function  TestMakeDefaultCACertifiatesPattern.

// Returns the default service settings in a map of key-value pairs.
func serviceSettings() SettingsMap {
	return map[string]any{
		"Kind":                 "station",
		"BindAddress":          DEFAULT_BIND_ADDRESS,
		"BindPort":             DEFAULT_BIND_PORT_NUMBER,
		"CACertPatterns":       makeDefaultCACertificatesPattern(),
		"DisableSubscriptions": false,
		"BufferSize":           DEFAULT_READ_BUFFER_SIZE,
		"MaxMessageSize":       DEFAULT_MAX_MESSAGE_SIZE,
		"MaxConcurrentStreams": DEFAULT_MAX_CONCURRENT_STREAMS,
		"NumStreamWorkers":     DEFAULT_NUM_STREAM_WORKERS,
	}

} // End of function  serviceSettings.

// Test makeDefaultServiceSettings function.
func TestMakeDefaultServiceSettings(t *testing.T) {
	expectation := serviceSettings()

	cfg := makeDefaultServiceSettings()

	if err := checkObject(cfg, expectation); err != nil {
		t.Error(err)
	}

} // End of function  TestMakeDefaultServiceSettings.

// Return config Test cases.
func configTestCases() []configTestCase {
	defaults := map[string]SettingsMap{
		"Settings": defaultSettings(),
		"Timeouts": timeoutSettings(),
		"Device":   deviceSettings(),
		"Service":  serviceSettings(),
	}

	telex := map[string]SettingsMap{
		"Settings": {
			"Name":        "WRU",
			"Debug":       false,
			"Version":     "0.0.7",
			"ValidateTLS": false,
			"Cert":        "",
			"Key":         "",
		},
		"Timeouts": timeoutSettings(),
		"Device": map[string]any{
			"Token":          "",
			"ServiceAddress": "127.1.2.3",
			"ServicePort":    9876,
			"ServiceCACert":  "",
			"RetryQueueSize": uint32(280),
		},
		"Service": serviceSettings(),
	}

	device := map[string]SettingsMap{
		"Settings": {
			"Name":        "dev-em",
			"Debug":       true,
			"Version":     "0.4.2",
			"ValidateTLS": true,
			"Cert":        "test/tls/device/telegraph-cert.pem",
			"Key":         "test/tls/device/telegraph-key.pem",
		},
		"Timeouts": timeoutSettings(),
		"Device": map[string]any{
			"Token":          "let me inside",
			"ServiceAddress": "127.0.0.1",
			"ServicePort":    9340,
			"ServiceCACert":  "test/tls/service/cacert.pem",
			"RetryQueueSize": uint32(2048),
		},
		"Service": serviceSettings(),
	}

	svcdev := deviceSettings()
	svcdev["ServiceAddress"] = "upstream.service.telegraph.local"

	service := map[string]SettingsMap{
		"Settings": {
			"Name":        "me:dev-ac",
			"Debug":       true,
			"Version":     "0.4.2",
			"ValidateTLS": true,
			"Cert":        "test/tls/service/bundle/service.pem",
			"Key":         "test/tls/service/bundle/service.pem",
		},
		"Timeouts": timeoutSettings(),
		"Device":   svcdev,
		"Service": {
			"Kind":        "station",
			"BindAddress": DEFAULT_BIND_ADDRESS,
			"BindPort":    DEFAULT_BIND_PORT_NUMBER,

			"CACertPatterns": CACertificatesPattern{
				Bootstrap: "test/tls/bootstrap/*-cacert.pem",
				Device:    "test/tls/device/*-cacert.pem",
			},

			"DisableSubscriptions": false,
			"BufferSize":           uint32(4194304),
			"MaxMessageSize":       uint32(4194304),
			"MaxConcurrentStreams": uint32(255),
			"NumStreamWorkers":     uint32(100),
		},
	}

	return []configTestCase{
		{
			name:        "missing file",
			namespace:   ENV_NAMESPACE,
			path:        "/tmp/path/to/missing/file/404.env",
			defaults:    defaults,
			expected:    defaults,
			errors:      true,
			nextensions: -1,
		},
		{
			name:        "invalid env file",
			namespace:   ENV_NAMESPACE,
			path:        filepath.Join("fixtures", "invalid.env"),
			defaults:    defaults,
			expected:    defaults,
			errors:      true,
			nextensions: -1,
		},
		{
			name:        "empty env file",
			namespace:   ENV_NAMESPACE,
			path:        filepath.Join("fixtures", "empty.env"),
			defaults:    defaults,
			expected:    defaults,
			errors:      false,
			nextensions: -1,
		},
		{
			name:        "test env file default namespace",
			namespace:   ENV_NAMESPACE,
			path:        filepath.Join("fixtures", "test.env"),
			defaults:    defaults,
			expected:    defaults,
			errors:      false,
			nextensions: 24,
		},
		{
			name:        "test env file telex namespace",
			namespace:   "TELEX",
			path:        filepath.Join("fixtures", "test.env"),
			defaults:    defaults,
			expected:    telex,
			errors:      false,
			nextensions: 18,
		},
		{
			name:        "device env file",
			namespace:   "GRPC_TELEGRAPH",
			path:        filepath.Join("../..", "config", "device.env"),
			defaults:    defaults,
			expected:    device,
			errors:      false,
			nextensions: 14,
		},
		{
			name:        "service env file",
			namespace:   "GRPC_TELEGRAPH",
			path:        filepath.Join("../..", "config", "service.env"),
			defaults:    defaults,
			expected:    service,
			errors:      false,
			nextensions: 16,
		},
	}

} // End of function  configTestCases.

// Test makeConfig function.
func TestMakeConfig(t *testing.T) {
	for _, step := range configTestCases() {
		cfg := makeConfig(step.namespace, step.path)

		fields := map[string]any{
			"Settings": cfg.Settings,
			"Timeouts": cfg.Timeouts,
			"Device":   cfg.Device,
			"Service":  cfg.Service,
		}

		for k, v := range fields {
			if expected, ok := step.defaults[k]; ok {
				if err := checkObject(v, expected); err != nil {
					t.Error(err)
				}
			}
		}

		if len(cfg.Extensions) != 0 {
			t.Errorf("%v expected no extensions, got %v", step.name,
				len(cfg.Extensions))
		}

		if len(cfg.tags) == 0 {
			t.Errorf("%v expected tags, found none", step.name)
		}
	}

} // End of function  TestMakeConfig.

// Check config fields.
func checkConfigFields(cfg *Config, expected map[string]SettingsMap) error {
	fields := map[string]any{
		"Settings": cfg.Settings,
		"Timeouts": cfg.Timeouts,
		"Device":   cfg.Device,
		"Service":  cfg.Service,
	}

	for k, v := range fields {
		if value, ok := expected[k]; ok {
			if err := checkObject(v, value); err != nil {
				return err
			}
		}
	}

	return nil

} // End of function  checkConfigFields.

// Test Config.Load
func TestConfigLoad(t *testing.T) {
	for _, step := range configTestCases() {
		cfg := makeConfig(step.namespace, step.path)

		if err := checkConfigFields(&cfg, step.defaults); err != nil {
			t.Errorf("%v error: %v", step.name, err)
		}

		if len(cfg.Extensions) != 0 {
			t.Errorf("%v expected no extensions, got %v", step.name,
				len(cfg.Extensions))
		}

		if len(cfg.tags) == 0 {
			t.Errorf("%v expected tags, found none", step.name)
		}

		err := cfg.Load()

		if step.errors {
			if err == nil {
				t.Errorf("%v load expected errors, got none", step.name)
			}

		} else if err != nil {
			t.Errorf("%v load expected no errors, got %v", step.name,
				err)
		}

		if err := checkConfigFields(&cfg, step.expected); err != nil {
			t.Errorf("%v error: %v", step.name, err)
		}

		if step.nextensions >= 0 {
			actual := len(cfg.Extensions)
			if actual != step.nextensions {
				t.Errorf("%v expected %v extensions, got %v", step.name,
					step.nextensions, actual)
			}
		}
	}

} // End of function  TestConfigLoad.

// Test NewConfig function.
func TestNewConfig(t *testing.T) {
	for _, step := range configTestCases() {
		cfg, err := NewConfig(step.namespace, step.path)
		if step.errors {
			if err == nil {
				t.Errorf("%v load expected errors, got none", step.name)
			}

			continue

		} else if err != nil {
			t.Errorf("%v load expected no errors, got %v", step.name,
				err)
		}

		if err := checkConfigFields(cfg, step.expected); err != nil {
			t.Errorf("%v error: %v", step.name, err)
		}

		if step.nextensions >= 0 {
			actual := len(cfg.Extensions)
			if actual != step.nextensions {
				t.Errorf("%v expected %v extensions, got %v", step.name,
					step.nextensions, actual)
			}
		}

		if len(cfg.tags) == 0 {
			t.Errorf("%v expected tags, found none", step.name)
		}
	}
} // End of function  TestNewConfig.

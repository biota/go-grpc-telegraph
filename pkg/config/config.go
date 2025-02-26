package config

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/biota/go-grpc-telegraph/pkg/util"
)

const (
	ENV_NAMESPACE = "GRPC_TELEGRAPH"

	// Default name.
	DEFAULT_NAME = "grpc-telegraph"

	// Default version.
	DEFAULT_VERSION = "0.4.2"

	// Default debug flag.
	DEFAULT_DEBUG_FLAG = false

	// Default device token.
	DEFAULT_DEVICE_TOKEN = ""

	// Default service type is station ... basically anything other than
	// station, means that we have another upstream to connect to.
	DEFAULT_SERVICE_TYPE = "station"

	// Default gRPC service local bind address and port number.
	DEFAULT_BIND_ADDRESS     = "127.0.0.1"
	DEFAULT_BIND_PORT_NUMBER = 9340

	// Default gRPC remote service address and port number.
	DEFAULT_SERVICE_ADDRESS     = "service.telegraph.local"
	DEFAULT_SERVICE_PORT_NUMBER = 9340

	// Default validate TLS configuration artifacts.
	DEFAULT_VALIDATE_TLS_CONFIG = true

	// Default device or service certificate and private key.
	DEFAULT_CERTIFICATE = ""
	DEFAULT_PRIVATE_KEY = ""

	// Default file pattern for bootstrap CA certificates.
	DEFAULT_BOOTSTRAP_CACERTS_PATTERN = ""

	// Default file pattern for device CA certificates.
	DEFAULT_DEVICE_CACERTS_PATTERN = ""

	// Default service CA certificate.
	DEFAULT_SERVICE_CACERT = ""

	// Default disable subscriptions.
	DEFAULT_DISABLE_SUBSCRIPTIONS = false

	// Default read buffer size and max message size - read in 1 go!
	DEFAULT_READ_BUFFER_SIZE = uint32(64 * 1024) // 64kb
	DEFAULT_MAX_MESSAGE_SIZE = uint32(64 * 1024) // 64kb

	// Default max concurrent streams and number of stream workers for
	// a service. Could probably leave at uint16 but making it uint32
	// for consistency (32-bit or bust!).
	DEFAULT_MAX_CONCURRENT_STREAMS = uint32(256)
	DEFAULT_NUM_STREAM_WORKERS     = uint32(8)

	// Max queue size for retries due to failures (example if service is
	// down - we can cache these many messages and resend them when we
	// regain connectivity). The rest we just drop on the floor.
	// Tune according to your available memory.
	//
	// Note: With the defaults this works out to max 64mb memory usage.
	//       1k messages x 64kb per message.
	//
	DEFAULT_RETRY_QUEUE_SIZE = uint32(1024)

	// Default Timeouts.
	DEFAULT_CONNECT_TIMEOUT    = time.Duration(20) * time.Second
	DEFAULT_SEND_TIMEOUT       = time.Duration(300) * time.Second
	DEFAULT_KEEP_ALIVE_TIMEOUT = time.Duration(300) * time.Second
	MIN_KEEP_ALIVE_TIMEOUT     = time.Duration(30) * time.Second

	// Default max subscription delay.
	DEFAULT_MAX_SUBSCRIPTION_DELAY = time.Duration(300) * time.Second

	// Struct env tag.
	STRUCT_ENV_TAG = "env"
)

// Common device and service settings.
type Settings struct {
	Name        string `env:"NAME"`
	Version     string `env:"VERSION"`
	Debug       bool   `env:"DEBUG"`
	ValidateTLS bool   `env:"VALIDATE_TLS_CONFIG"`
	Cert        string `env:"CERT"`
	Key         string `env:"KEY"`
}

// Timeout settings.
type TimeoutSettings struct {
	Connect              time.Duration `env:"CONNECT_TIMEOUT"`
	Send                 time.Duration `env:"SEND_TIMEOUT"`
	KeepAlive            time.Duration `env:"KEEP_ALIVE_TIMEOUT"`
	MaxSubscriptionDelay time.Duration `env:"MAX_SUBSCRIPTION_DELAY"`
}

// Device settings.
type DeviceSettings struct {
	Token          string `env:"TOKEN"`
	ServiceAddress string `env:"SERVICE_ADDRESS"`
	ServicePort    int    `env:"SERVICE_PORT"`
	ServiceCACert  string `env:"SERVICE_CACERT"`
	RetryQueueSize uint32 `env:"RETRY_QUEUE_SIZE"`
}

// CA certificates pattern for bootstrap and device CAs.
type CACertificatesPattern struct {
	Bootstrap string `env:"BOOTSTRAP_CACERTS_PATTERN"`
	Device    string `env:"DEVICE_CACERTS_PATTERN"`
}

// Service settings.
type ServiceSettings struct {
	Kind        string `env:"TYPE"`
	BindAddress string `env:"BIND_ADDRESS"`
	BindPort    int    `env:"BIND_PORT"`

	CACertPatterns CACertificatesPattern

	DisableSubscriptions bool   `env:"DISABLE_SUBSCRIPTIONS"`
	BufferSize           uint32 `env:"BUFFER_SIZE"`
	MaxMessageSize       uint32 `env:"MAX_MESSAGE_SIZE"`
	MaxConcurrentStreams uint32 `env:"MAX_STREAMS"`
	NumStreamWorkers     uint32 `env:"NUM_STREAM_WORKERS"`
}

// Telegraph configuration loaded from defaults/environment/settings file.
type Config struct {
	Namespace  string
	File       string
	Settings   Settings
	Timeouts   TimeoutSettings
	Device     DeviceSettings
	Service    ServiceSettings
	Extensions map[string]string

	tags map[string][]string
}

// Make default settings.
func makeDefaultSettings() Settings {
	return Settings{
		Name:        DEFAULT_NAME,
		Version:     DEFAULT_VERSION,
		Debug:       DEFAULT_DEBUG_FLAG,
		ValidateTLS: DEFAULT_VALIDATE_TLS_CONFIG,
		Cert:        DEFAULT_CERTIFICATE,
		Key:         DEFAULT_PRIVATE_KEY,
	}

} //  End of function  makeDefaultSettings.

// Make default timeout settings.
func makeDefaultTimeoutSettings() TimeoutSettings {
	return TimeoutSettings{
		Connect:              DEFAULT_CONNECT_TIMEOUT,
		Send:                 DEFAULT_SEND_TIMEOUT,
		KeepAlive:            DEFAULT_KEEP_ALIVE_TIMEOUT,
		MaxSubscriptionDelay: DEFAULT_MAX_SUBSCRIPTION_DELAY,
	}

} //  End of function  makeDefaultTimeoutSettings.

// Make default device settings.
func makeDefaultDeviceSettings() DeviceSettings {
	return DeviceSettings{
		Token:          DEFAULT_DEVICE_TOKEN,
		ServiceAddress: DEFAULT_SERVICE_ADDRESS,
		ServicePort:    DEFAULT_SERVICE_PORT_NUMBER,
		ServiceCACert:  DEFAULT_SERVICE_CACERT,
		RetryQueueSize: DEFAULT_RETRY_QUEUE_SIZE,
	}

} //  End of function  makeDefaultDeviceSettings.

// Make default CA cerfificates pattern.
func makeDefaultCACertificatesPattern() CACertificatesPattern {
	return CACertificatesPattern{
		Bootstrap: DEFAULT_BOOTSTRAP_CACERTS_PATTERN,
		Device:    DEFAULT_DEVICE_CACERTS_PATTERN,
	}

} // End of function  makeDefaultCACertificatesPattern.

// Make default service settings.
func makeDefaultServiceSettings() ServiceSettings {
	caCertPatterns := makeDefaultCACertificatesPattern()

	return ServiceSettings{
		Kind:                 DEFAULT_SERVICE_TYPE,
		BindAddress:          DEFAULT_BIND_ADDRESS,
		BindPort:             DEFAULT_BIND_PORT_NUMBER,
		CACertPatterns:       caCertPatterns,
		DisableSubscriptions: DEFAULT_DISABLE_SUBSCRIPTIONS,
		BufferSize:           DEFAULT_READ_BUFFER_SIZE,
		MaxMessageSize:       DEFAULT_MAX_MESSAGE_SIZE,
		MaxConcurrentStreams: DEFAULT_MAX_CONCURRENT_STREAMS,
		NumStreamWorkers:     DEFAULT_NUM_STREAM_WORKERS,
	}

} //  End of function  makeDefaultServiceSettings.

// Make configuration.
func makeConfig(namespace, path string) Config {
	cfg := Config{
		Namespace:  namespace,
		File:       path,
		Settings:   makeDefaultSettings(),
		Timeouts:   makeDefaultTimeoutSettings(),
		Device:     makeDefaultDeviceSettings(),
		Service:    makeDefaultServiceSettings(),
		Extensions: make(map[string]string),

		tags: make(map[string][]string),
	}

	// Build the list of tags for speeding up setting updates.
	cfg.tags["Settings"] = util.StructTags(cfg.Settings, STRUCT_ENV_TAG)
	cfg.tags["Timeouts"] = util.StructTags(cfg.Timeouts, STRUCT_ENV_TAG)
	cfg.tags["Device"] = util.StructTags(cfg.Device, STRUCT_ENV_TAG)
	cfg.tags["Service"] = util.StructTags(cfg.Service, STRUCT_ENV_TAG)

	return cfg

} // End of function  makeConfig.

// Update a general/common configuration setting.
func (c *Config) updateSettings(name, value string) error {
	switch strings.ToUpper(name) {
	case "NAME":
		c.Settings.Name = value

	case "VERSION":
		c.Settings.Version = value

	case "DEBUG":
		if v, err := util.ToBoolean(value); err == nil {
			c.Settings.Debug = v
		} else {
			return err
		}

	case "VALIDATE_TLS_CONFIG":
		if v, err := util.ToBoolean(value); err == nil {
			c.Settings.ValidateTLS = v
		} else {
			return err
		}

	case "CERT":
		c.Settings.Cert = value

	case "KEY":
		c.Settings.Key = value
	}

	return nil

} // End of  Config.updateSettings

// Update a configuration timeout setting.
func (c *Config) updateTimeouts(name, value string) error {
	v, err := util.ToTimeDuration(value)
	if err != nil {
		return err
	}

	switch strings.ToUpper(name) {
	case "CONNECT":
		c.Timeouts.Connect = v

	case "SEND":
		c.Timeouts.Send = v

	case "KEEP_ALIVE":
		c.Timeouts.KeepAlive = max(v, MIN_KEEP_ALIVE_TIMEOUT)

	case "MAX_SUBSCRIPTION_DELAY":
		c.Timeouts.MaxSubscriptionDelay = v
	}

	return nil

} // End of  Config.updateTimeouts

// Update a device setting.
func (c *Config) updateDeviceSettings(name, value string) error {
	switch strings.ToUpper(name) {
	case "TOKEN":
		c.Device.Token = value

	case "SERVICE_ADDRESS":
		c.Device.ServiceAddress = value

	case "SERVICE_PORT":
		if v, err := util.ToInteger(value); err == nil {
			c.Device.ServicePort = v
		} else {
			return err
		}

	case "SERVICE_CACERT":
		c.Device.ServiceCACert = value

	case "RETRY_QUEUE_SIZE":
		if v, err := util.ToUnsignedInt32(value); err == nil {
			c.Device.RetryQueueSize = v
		} else {
			return err
		}
	}

	return nil

} // End of  Config.updateDeviceSettings

// Update a service setting.
func (c *Config) updateServiceSettings(name, value string) error {
	switch strings.ToUpper(name) {
	case "TYPE":
		c.Service.Kind = value

	case "BIND_ADDRESS":
		c.Service.BindAddress = value

	case "BIND_PORT":
		if v, err := util.ToInteger(value); err == nil {
			c.Service.BindPort = v
		} else {
			return err
		}

	case "BOOTSTRAP_CACERTS_PATTERN":
		c.Service.CACertPatterns.Bootstrap = value

	case "DEVICE_CACERTS_PATTERN":
		c.Service.CACertPatterns.Device = value

	case "DISABLE_SUBSCRIPTIONS":
		if v, err := util.ToBoolean(value); err == nil {
			c.Service.DisableSubscriptions = v
		} else {
			return err
		}

	case "BUFFER_SIZE":
		if v, err := util.ToUnsignedInt32(value); err == nil {
			c.Service.BufferSize = v
		} else {
			return err
		}

	case "MAX_MESSAGE_SIZE":
		if v, err := util.ToUnsignedInt32(value); err == nil {
			c.Service.MaxMessageSize = v
		} else {
			return err
		}

	case "MAX_STREAMS":
		if v, err := util.ToUnsignedInt32(value); err == nil {
			c.Service.MaxConcurrentStreams = v
		} else {
			return err
		}

	case "NUM_STREAM_WORKERS":
		if v, err := util.ToUnsignedInt32(value); err == nil {
			c.Service.NumStreamWorkers = v
		} else {
			return err
		}
	}

	return nil

} // End of  Config.updateServiceSettings

// Update value for a configuration setting.
func (c *Config) update(name, value string) error {
	prefix := c.Namespace + "_"
	suffix := name

	if strings.HasPrefix(name, prefix) {
		suffix = strings.TrimPrefix(name, prefix)
	}

	if slices.Contains(c.tags["Settings"], suffix) {
		return c.updateSettings(suffix, value)
	}

	if slices.Contains(c.tags["Timeouts"], suffix) {
		return c.updateTimeouts(suffix, value)
	}

	if slices.Contains(c.tags["Device"], suffix) {
		return c.updateDeviceSettings(suffix, value)
	}

	if slices.Contains(c.tags["Service"], suffix) {
		return c.updateServiceSettings(suffix, value)
	}

	// Any other setting just goes directly into the extensions map.
	c.Extensions[name] = value

	return nil

} // End of  Config.update

// Load config from environment values.
func (c *Config) loadEnv() error {
	prefix := c.Namespace + "_"

	for _, e := range os.Environ() {
		pieces := strings.SplitN(e, "=", 2)
		if len(pieces) != 2 {
			return fmt.Errorf("invalid env entry: %v", e)
		}

		if !strings.HasPrefix(pieces[0], prefix) {
			// Skip non-namespaced variables.
			continue
		}

		if err := c.update(pieces[0], pieces[1]); err != nil {
			slog.Error("loading env variable", "error", err,
				e, pieces)
			return err
		}
	}

	return nil

} // End of  Config.loadEnv

// Load config from an environment settings file.
func (c *Config) loadSettings() error {
	if len(c.File) == 0 {
		// No settings file, nothing to do.
		return nil
	}

	cfg, err := LoadSettings(c.File)
	if err != nil {
		return err
	}

	for k, v := range cfg {
		if err := c.update(k, v); err != nil {
			slog.Error("loading setting", "error", err, k, v)
			return err
		}
	}

	return nil

} // End of  Config.loadSettings.

// Load configuration.
func (c *Config) Load() error {
	// Load environment values first, then load the env settings file.
	// That way the settings file entry takes precedence.
	if err := c.loadEnv(); err != nil {
		return err
	}

	return c.loadSettings()

} // End of  Config.Load

// Return a new config instance.
func NewConfig(namespace, file string) (*Config, error) {
	cfg := makeConfig(namespace, file)
	err := cfg.Load()

	return &cfg, err

} // End of function  NewConfig.

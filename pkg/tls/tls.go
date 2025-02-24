package tls

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	CERTIFICATE     = "CERTIFICATE"
	EC_PRIVATE_KEY  = "EC PRIVATE KEY"
	RSA_PRIVATE_KEY = "RSA PRIVATE KEY"
	PRIVATE_KEY     = "PRIVATE KEY"
)

// Errors loading secure assets.
type LoadAssetErrors struct {
	Errors []error
}

// Returns combined load asset errors - implements `error` interface.
func (e *LoadAssetErrors) Error() string {
	messages := []string{}
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}

	return "[" + strings.Join(messages, ", ") + "]"

} //  End of  LoadAssetErrors.Error

// Finds matching PEM blocks for a specific type.
func findBlocks(path, kind string) ([]*pem.Block, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	matched := make([]*pem.Block, 0)

	remaining := data
	for {
		block, rest := pem.Decode(remaining)
		if block == nil {
			break
		}

		remaining = rest

		if block.Type == kind {
			matched = append(matched, block)
		}
	}

	if len(matched) > 0 {
		return matched, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("no matching %v blocks", kind))

} //  End of function  findBlocks.

// Loads all PEM format certificates from a file.
func LoadCertificates(path string) ([]*x509.Certificate, error) {
	certs := []*x509.Certificate{}
	errs := []error{}

	blocks, err := findBlocks(path, CERTIFICATE)
	if err != nil {
		errs = append(errs, err)
		return nil, &LoadAssetErrors{Errors: errs}
	}

	for _, block := range blocks {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err == nil {
			certs = append(certs, cert)
		} else {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return certs, &LoadAssetErrors{Errors: errs}
	}

	return certs, nil

} //  End of function  LoadCertificates.

// Validate all PEM format certificates in a file.
func ValidateCertificates(path string) error {
	_, err := LoadCertificates(path)
	return err

} //  End of function  ValidateCertificates.

// Loads all PEM format CA certificates from a file.
func LoadCACerts(path string) ([]*x509.Certificate, error) {
	return LoadCertificates(path)

} //  End of function  LoadCACerts.

// Validate all PEM format CA certificates in a file.
func ValidateCACerts(path string) error {
	_, err := LoadCACerts(path)
	return err

} //  End of function  ValidateCACerts.

// Parse a private key from a PEM block.
func parsePrivateKey(block *pem.Block) (any, error) {
	switch block.Type {
	case EC_PRIVATE_KEY:
		return x509.ParseECPrivateKey(block.Bytes)

	case RSA_PRIVATE_KEY:
		return x509.ParsePKCS1PrivateKey(block.Bytes)

	case PRIVATE_KEY:
		return x509.ParsePKCS8PrivateKey(block.Bytes)
	}

	return nil, fmt.Errorf("unsupported PEM type: %v", block.Type)

} //  End of function  parsePrivateKey.

// Loads all PEM format private keys from a file.
func LoadPrivateKeys(path string) ([]any, error) {
	privateKeys := []any{}
	errs := []error{}

	blocks, err := findBlocks(path, PRIVATE_KEY)
	if err != nil {
		return nil, err
	}

	for _, block := range blocks {
		zkey, err := parsePrivateKey(block)
		if err == nil {
			privateKeys = append(privateKeys, zkey)
		} else {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return privateKeys, &LoadAssetErrors{Errors: errs}
	}

	return privateKeys, nil

} //  End of function  LoadPrivateKeys.

// Validate all PEM format private keys in a file.
func ValidatePrivateKeys(path string) error {
	_, err := LoadPrivateKeys(path)
	return err

} //  End of function  ValidatePrivateKeys.

// Load Certificate and Key pair from given paths.
func LoadCertKeyPair(certPath, keyPath string) (tls.Certificate, error) {
	return tls.LoadX509KeyPair(certPath, keyPath)

} //  End of function  LoadCertKeyPair.

// Validate a Certificate and Key pair.
func ValidateCertKeyPair(certPath, keyPath string) error {
	_, err := LoadCertKeyPair(certPath, keyPath)
	return err

} //  End of function  ValidateCertKeyPair.

// Creates a [CA] certificate pool loaded with the valid CA certificates.
// Error indicates one or more assets failed to be loaded but the pool
// still contains all the valid CA certificates.
func CreateCACertPool(caCertPaths []string) (*x509.CertPool, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	errs := []error{}

	// Collect the errors ... and load valid certificates.
	for _, zpath := range caCertPaths {
		cacerts, err := LoadCACerts(zpath)
		if err != nil {
			log.Printf("ERROR: Adding CA %v to cert pool: %v",
				zpath, err)
			errs = append(errs, err)
			continue
		}

		// We could done the load in one shot by saving 'em off to
		// a local list ... but one bad "apple" shouldn't spoil
		// the barrel ..  as long you don't add it in!
		for _, cert := range cacerts {
			pool.AddCert(cert)
		}
	}

	if len(errs) > 0 {
		return pool, &LoadAssetErrors{Errors: errs}
	}

	return pool, nil

} // End of function  CreateCACertPool.

// Create secure device TLS config.
// Normally, this is the client side config but in a tower workflow, it
// would be the "service in the middle" attack!
func DeviceConfig(certPath, keyPath string,
	caCertPaths []string) (*tls.Config, error) {

	pool, err := CreateCACertPool(caCertPaths)
	if err != nil {
		return nil, err
	}

	certs := []tls.Certificate{}
	if len(certPath) > 0 || len(keyPath) > 0 {
		cert, err := LoadCertKeyPair(certPath, keyPath)
		if err != nil {
			return nil, err
		}

		certs = append(certs, cert)
	}

	return &tls.Config{Certificates: certs, RootCAs: pool}, nil

} // End of function DeviceConfig.

// Create secure service config.
func ServiceConfig(certPath, keyPath string, clientCACertPaths []string,
	clientAuth tls.ClientAuthType) (*tls.Config, error) {

	pool, err := CreateCACertPool(clientCACertPaths)
	if err != nil {
		return nil, err
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{},
		ClientAuth:   clientAuth,
		ClientCAs:    pool,
	}

	if len(certPath) > 0 || len(keyPath) > 0 {
		cert, err := LoadCertKeyPair(certPath, keyPath)
		if err != nil {
			return nil, err
		}

		cfg.Certificates = append(cfg.Certificates, cert)
	}

	return cfg, nil

} // End of function  ServiceConfig.

package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	DEVICE            = "device"
	SERVICE           = "service"
	BOOTSTRAP         = "bootstrap"
	DEVICE_BUNDLE     = "device/bundle"
	SERVICE_BUNDLE    = "service/bundle"
	BOOTSTRAP_BUNDLE  = "bootstrap/bundle"
	TEST_FIXTURES_DIR = "/tmp/go-grpc-telegraph/test"
)

type tlsTestCase struct {
	name     string
	certpath string
	keypath  string
	ncerts   int
	nkeys    int
	loadErr  bool
}

type tlsCATestCase struct {
	name      string
	cacerts   []string
	loadError bool
}

type LoadCertsFunc func(path string) ([]*x509.Certificate, error)
type ValidateCertsFunc func(path string) error

// Test LoadAssetErrors
func TestLoadAssetErrors(t *testing.T) {
	tests := []struct {
		name   string
		errors []error
	}{
		{
			name:   "empty",
			errors: []error{},
		},
		{
			name: "nil",
		},
		{
			name:   "one",
			errors: []error{fmt.Errorf("one error")},
		},
		{
			name: "dos",
			errors: []error{fmt.Errorf("uno"),
				fmt.Errorf("dos")},
		},
		{
			name: "trois",
			errors: []error{fmt.Errorf("une"),
				fmt.Errorf("deux"),
				fmt.Errorf("trois"),
			},
		},
		{
			name: "many",
			errors: []error{fmt.Errorf("ek"), fmt.Errorf("do"),
				fmt.Errorf("teen"), fmt.Errorf("char"),
				fmt.Errorf("paanch"),
			},
		},
	}

	for _, step := range tests {
		laerr := &LoadAssetErrors{Errors: step.errors}
		value := reflect.ValueOf(laerr)
		if _, ok := value.Interface().(error); !ok {
			t.Errorf("expected %v to implement error", laerr)
		}
	}

} //  End of function  TestLoadAssetErrors.

// Returns name of a function.
func functionName(fn interface{}) string {
	fqname := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	parts := strings.Split(fqname, ".")
	return parts[len(parts)-1]

} //  End of  functionName

// Get path to a tls config pem file.
func getPath(where, name string) string {
	repodir := "../.."

	if zdir, err := os.Getwd(); err == nil {
		zpath, err := filepath.Abs(filepath.Join(zdir, "..", ".."))
		if err == nil {
			repodir = zpath
		}
	}

	return filepath.Join(repodir, "config/test/tls", where, name)

} //  End of  getPath

func createMultiPEMFile(fileName string, paths ...string) string {
	if err := os.MkdirAll(TEST_FIXTURES_DIR, 0750); err != nil {
		log.Fatal(err)
	}

	content := []byte{}
	tempFileName := filepath.Join(TEST_FIXTURES_DIR, fileName)

	for _, zpath := range paths {
		if data, err := os.ReadFile(zpath); err == nil {
			content = append(content[:], data[:]...)
		}
	}

	err := os.WriteFile(tempFileName, content, 0777)
	if err != nil {
		fmt.Printf("Error writing %v: %v\n", tempFileName, err)
	}

	return tempFileName

} //  End of createMultiPEMFile

// Generate TLS test cases.
func generateTestCases() []tlsTestCase {
	emptyPEM := filepath.Join(TEST_FIXTURES_DIR, "empty.pem")
	if err := os.WriteFile(emptyPEM, []byte{}, 0777); err != nil {
		fmt.Printf("Error writing %v: %v\n", emptyPEM, err)
	}

	badPEM := filepath.Join(TEST_FIXTURES_DIR, "bad.pem")
	if err := os.WriteFile(badPEM, []byte("baddy"), 0777); err != nil {
		fmt.Printf("Error writing %v: %v\n", badPEM, err)
	}

	c5certs := createMultiPEMFile("c5certs.pem",
		getPath(DEVICE, "alchemy-cert.pem"),
		getPath(DEVICE, "communique-cert.pem"),
		getPath(DEVICE, "industrial-disease-cert.pem"),
		getPath(DEVICE, "news-cert.pem"),
		getPath(DEVICE, "telegraph-cert.pem"))

	c5cacerts := createMultiPEMFile("c5cacerts.pem",
		getPath(DEVICE, "alchemy-cacert.pem"),
		getPath(DEVICE, "communique-cacert.pem"),
		getPath(DEVICE, "industrial-disease-cacert.pem"),
		getPath(DEVICE, "news-cacert.pem"),
		getPath(DEVICE, "telegraph-cacert.pem"))

	cscerts := createMultiPEMFile("cscerts.pem",
		getPath(DEVICE, "alchemy-cert.pem"),
		getPath(BOOTSTRAP, "bootstrap-cert.pem"),
		getPath(SERVICE, "cert.pem"))

	cscertsca := createMultiPEMFile("cscertsca.pem",
		getPath(DEVICE, "telegraph-cert.pem"),
		getPath(DEVICE, "news-cacert.pem"),
		getPath(SERVICE, "cert.pem"),
		getPath(SERVICE, "cacert.pem"))

	combocerts := createMultiPEMFile("combocerts.pem",
		getPath(DEVICE, "alchemy-cert.pem"),
		getPath(DEVICE, "alchemy-cacert.pem"),
		getPath(DEVICE, "communique-cert.pem"),
		getPath(DEVICE, "communique-cacert.pem"),
		getPath(DEVICE, "industrial-disease-cert.pem"),
		getPath(DEVICE, "industrial-disease-cacert.pem"),
		getPath(DEVICE, "news-cert.pem"),
		getPath(DEVICE, "news-cacert.pem"),
		getPath(DEVICE, "telegraph-cert.pem"),
		getPath(SERVICE, "cert.pem"),
		getPath(SERVICE, "cacert.pem"),
		getPath(BOOTSTRAP, "bootstrap-cert.pem"))

	c5keys := createMultiPEMFile("c5keys.pem",
		getPath(DEVICE, "alchemy-key.pem"),
		getPath(DEVICE, "communique-key.pem"),
		getPath(DEVICE, "industrial-disease-key.pem"),
		getPath(DEVICE, "news-key.pem"),
		getPath(DEVICE, "telegraph-key.pem"))

	cskeys := createMultiPEMFile("cskeys.pem",
		getPath(DEVICE, "news-key.pem"),
		getPath(SERVICE, "key.pem"),
		getPath(BOOTSTRAP, "bootstrap-key.pem"))

	comboall := createMultiPEMFile("combokeys.pem",
		getPath(DEVICE, "alchemy-cacert.pem"),
		getPath(DEVICE, "alchemy-cakey.pem"),
		getPath(DEVICE, "alchemy-cert.pem"),
		getPath(DEVICE, "alchemy-key.pem"),
		getPath(DEVICE, "alchemy-csr.pem"),
		getPath(DEVICE, "communique-cert.pem"),
		getPath(DEVICE, "news-key.pem"),
		getPath(DEVICE, "news-cert.pem"),
		getPath(DEVICE, "news-csr.pem"),
		getPath(DEVICE, "telegraph-cacert.pem"),
		getPath(DEVICE_BUNDLE, "ca-alchemy.pem"),
		getPath(DEVICE_BUNDLE, "ca-news.pem"),
		getPath(DEVICE_BUNDLE, "device-communique.pem"),
		getPath(DEVICE_BUNDLE, "device-news.pem"),
		getPath(DEVICE_BUNDLE, "device-telegraph.pem"),
		getPath(SERVICE, "cert.pem"),
		getPath(SERVICE, "key.pem"),
		getPath(SERVICE, "cacert.pem"),
		getPath(SERVICE, "cakey.pem"),
		getPath(SERVICE_BUNDLE, "ca.pem"),
		getPath(BOOTSTRAP, "bootstrap-cert.pem"),
		getPath(BOOTSTRAP, "bootstrap-key.pem"),
		getPath(BOOTSTRAP, "bootstrap-cacert.pem"),
		getPath(BOOTSTRAP, "bootstrap-cakey.pem"),
		getPath(BOOTSTRAP_BUNDLE, "ca-bootstrap.pem"))

	badcertkey := createMultiPEMFile("badcertkey.pem",
		getPath(DEVICE, "alchemy-cert.pem"),
		getPath(DEVICE, "news-cakey.pem"))

	badcombo := createMultiPEMFile("badcertkeycombo.pem",
		getPath(DEVICE, "alchemy-cert.pem"),
		getPath(DEVICE, "communique-cacert.pem"),
		getPath(DEVICE_BUNDLE, "ca-news.pem"),
		getPath(SERVICE, "cacert.pem"),
		getPath(DEVICE, "news-key.pem"),
		getPath(DEVICE, "telegraph-cakey.pem"),
		getPath(SERVICE, "key.pem"),
		getPath(BOOTSTRAP, "bootstrap-cakey.pem"))

	fixtures := []tlsTestCase{
		{
			name:     "Single device certificate",
			certpath: getPath(DEVICE, "alchemy-cert.pem"),
			keypath:  getPath(DEVICE, "alchemy-cert.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Another device certificate",
			certpath: getPath(DEVICE, "news-cert.pem"),
			keypath:  getPath(DEVICE, "news-cert.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Single device CA certificate",
			certpath: getPath(DEVICE, "alchemy-cacert.pem"),
			keypath:  getPath(DEVICE, "alchemy-cacert.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Another device CA certificate",
			certpath: getPath(DEVICE, "news-cacert.pem"),
			keypath:  getPath(DEVICE, "news-cacert.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name: "Bundled device certificate and key",
			certpath: getPath(DEVICE_BUNDLE,
				"device-alchemy.pem"),
			keypath: getPath(DEVICE_BUNDLE,
				"device-alchemy.pem"),
			ncerts:  1,
			nkeys:   1,
			loadErr: false,
		},
		{
			name:     "Another bundled device certificate+key",
			certpath: getPath(DEVICE_BUNDLE, "device-news.pem"),
			keypath:  getPath(DEVICE_BUNDLE, "device-news.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "Bundled device CA certificate",
			certpath: getPath(DEVICE_BUNDLE, "ca-alchemy.pem"),
			keypath:  getPath(DEVICE_BUNDLE, "ca-alchemy.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Another bundled device CA cert",
			certpath: getPath(DEVICE_BUNDLE, "ca-news.pem"),
			keypath:  getPath(DEVICE_BUNDLE, "ca-news.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Service certificate",
			certpath: getPath(SERVICE, "cert.pem"),
			keypath:  getPath(SERVICE, "cert.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Service CA certificate",
			certpath: getPath(SERVICE, "cacert.pem"),
			keypath:  getPath(SERVICE, "cacert.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "bundled service certificate and key",
			certpath: getPath(SERVICE_BUNDLE, "service.pem"),
			keypath:  getPath(SERVICE_BUNDLE, "service.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "bundled service CA certificate",
			certpath: getPath(SERVICE_BUNDLE, "ca.pem"),
			keypath:  getPath(SERVICE_BUNDLE, "ca.pem"),
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Single device key",
			certpath: getPath(DEVICE, "alchemy-key.pem"),
			keypath:  getPath(DEVICE, "alchemy-key.pem"),
			ncerts:   0,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "Single device CA key",
			certpath: getPath(DEVICE, "alchemy-cakey.pem"),
			keypath:  getPath(DEVICE, "alchemy-cakey.pem"),
			ncerts:   0,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "Single device CSR",
			certpath: getPath(DEVICE, "alchemy-csr.pem"),
			keypath:  getPath(DEVICE, "alchemy-csr.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Service key",
			certpath: getPath(SERVICE, "key.pem"),
			keypath:  getPath(SERVICE, "key.pem"),
			ncerts:   0,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "Service CA key",
			certpath: getPath(SERVICE, "cakey.pem"),
			keypath:  getPath(SERVICE, "cakey.pem"),
			ncerts:   0,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "Service CSR",
			certpath: getPath(SERVICE, "csr.pem"),
			keypath:  getPath(SERVICE, "csr.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Bootstrap key",
			certpath: getPath(BOOTSTRAP, "bootstrap-key.pem"),
			keypath:  getPath(BOOTSTRAP, "bootstrap-key.pem"),
			ncerts:   0,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name: "Bootstrap CA key",
			certpath: getPath(BOOTSTRAP,
				"bootstrap-cakey.pem"),
			keypath: getPath(BOOTSTRAP,
				"bootstrap-cakey.pem"),
			ncerts:  0,
			nkeys:   1,
			loadErr: true,
		},
		{
			name:     "Bootstrap CSR",
			certpath: getPath(BOOTSTRAP, "bootstrap-csr.pem"),
			keypath:  getPath(BOOTSTRAP, "bootstrap-csr.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Missing cert",
			certpath: getPath(DEVICE, "missing-cert.pem"),
			keypath:  getPath(DEVICE, "missing-cert.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Bad key",
			certpath: getPath(DEVICE, "bad-key.pem"),
			keypath:  getPath(DEVICE, "bad-key.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Invalid csr",
			certpath: getPath(DEVICE, "invalid-csr.pem"),
			keypath:  getPath(DEVICE, "invalid-csr.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Empty PEM file",
			certpath: emptyPEM,
			keypath:  emptyPEM,
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "Bad PEM file",
			certpath: badPEM,
			keypath:  badPEM,
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "5 device certs PEM file",
			certpath: c5certs,
			keypath:  c5certs,
			ncerts:   5,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "5 device CA certs PEM file",
			certpath: c5cacerts,
			keypath:  c5cacerts,
			ncerts:   5,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "device+service certs PEM file",
			certpath: cscerts,
			keypath:  cscerts,
			ncerts:   3,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "device+service+ca certs PEM file",
			certpath: cscertsca,
			keypath:  cscertsca,
			ncerts:   4,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "combo certs PEM file",
			certpath: combocerts,
			keypath:  combocerts,
			ncerts:   12,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "5 device keys PEM file",
			certpath: c5keys,
			keypath:  c5keys,
			ncerts:   0,
			nkeys:    5,
			loadErr:  true,
		},
		{
			name:     "device+service keys PEM file",
			certpath: cskeys,
			keypath:  cskeys,
			ncerts:   0,
			nkeys:    3,
			loadErr:  true,
		},
		{
			name:     "combo keys,certs,cacerts PEM file",
			certpath: comboall,
			keypath:  comboall,
			ncerts:   16,
			nkeys:    10,
			loadErr:  false,
		},
		{
			name:     "only certificate PEM file",
			certpath: getPath(DEVICE, "news-cert.pem"),
			keypath:  "",
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "only CA certificate PEM file",
			certpath: getPath(DEVICE, "news-cacert.pem"),
			keypath:  "",
			ncerts:   1,
			nkeys:    0,
			loadErr:  true,
		},
		{
			name:     "only private key PEM file",
			certpath: "",
			keypath:  getPath(DEVICE, "alchemy-key.pem"),
			ncerts:   0,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name: "valid certificate key pair",
			certpath: getPath(DEVICE_BUNDLE,
				"device-news.pem"),
			keypath: getPath(DEVICE_BUNDLE,
				"device-news.pem"),
			ncerts:  1,
			nkeys:   1,
			loadErr: false,
		},
		{
			name: "another valid certificate key pair",
			certpath: getPath(DEVICE_BUNDLE,
				"device-alchemy.pem"),
			keypath: getPath(DEVICE_BUNDLE,
				"device-alchemy.pem"),
			ncerts:  1,
			nkeys:   1,
			loadErr: false,
		},
		{
			name:     "Valid service certificate key pair",
			certpath: getPath(SERVICE_BUNDLE, "service.pem"),
			keypath:  getPath(SERVICE_BUNDLE, "service.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name: "Valid bootstrap certificate key pair",
			certpath: getPath(BOOTSTRAP_BUNDLE,
				"device-bootstrap.pem"),
			keypath: getPath(BOOTSTRAP_BUNDLE,
				"device-bootstrap.pem"),
			ncerts:  1,
			nkeys:   1,
			loadErr: false,
		},
		{
			name:     "Bad certificate key pair",
			certpath: badcertkey,
			keypath:  badcertkey,
			ncerts:   1,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "Bad certificate key pair combo",
			certpath: badcombo,
			keypath:  badcombo,
			ncerts:   4,
			nkeys:    4,
			loadErr:  true,
		},
		{
			name:     "Bad device certificate key pair",
			certpath: getPath(DEVICE, "alchemy-cert.pem"),
			keypath:  getPath(SERVICE, "key.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "Good device certificate key pair",
			certpath: getPath(DEVICE, "news-cert.pem"),
			keypath:  getPath(DEVICE, "news-key.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "another good device cert key pair",
			certpath: getPath(DEVICE, "telegraph-cert.pem"),
			keypath:  getPath(DEVICE, "telegraph-key.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "Good service certificate key pair",
			certpath: getPath(SERVICE, "cert.pem"),
			keypath:  getPath(SERVICE, "key.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "Good bootstrap certificate key pair",
			certpath: getPath(BOOTSTRAP, "bootstrap-cert.pem"),
			keypath:  getPath(BOOTSTRAP, "bootstrap-key.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "bundled device certificate+key pair",
			certpath: getPath(DEVICE_BUNDLE, "device-news.pem"),
			keypath:  getPath(DEVICE, "news-key.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  false,
		},
		{
			name:     "device certificate+key bundle pair",
			certpath: getPath(DEVICE, "telegraph-cert.pem"),
			keypath:  getPath(DEVICE_BUNDLE, "device-news.pem"),
			ncerts:   1,
			nkeys:    1,
			loadErr:  true,
		},
		{
			name:     "misconfigured certificate+key pair",
			certpath: getPath(DEVICE, "news-key.pem"),
			keypath:  getPath(DEVICE, "news-cert.pem"),
			ncerts:   0,
			nkeys:    0,
			loadErr:  true,
		},
	}

	return fixtures

} //  End of generateTestCases

// Check load certificates function.
func checkLoadCertsFunc(fn LoadCertsFunc, tc tlsTestCase) error {
	name := functionName(fn)

	fail := func(unit, msg string, err error) error {
		if err == nil {
			return fmt.Errorf("%v unit %v %v", name, unit, msg)
		}

		return fmt.Errorf("%v error: %w", msg, err)
	}

	certs, err := fn(tc.certpath)

	if tc.ncerts > 0 {
		if err != nil {
			return fail(tc.name, "unexpected", err)
		}

		ncerts := 0
		if certs != nil {
			ncerts = len(certs)
		}

		if ncerts != tc.ncerts {
			msg := fmt.Sprintf("expected %v certs, got %v",
				tc.ncerts, ncerts)
			return fail(tc.name, msg, nil)
		} else {
			fmt.Printf("test %v unit %v matched %v certs\n",
				name, tc.name, tc.ncerts)
		}

		return nil
	}

	// we didn't get any certs, so expect errors.
	if err != nil {
		return nil
	}

	return fail(tc.name, "expected cert error", err)

} //  End of function  checkLoadCertsFunc.

// Test LoadCertificates function.
func TestLoadCertificates(t *testing.T) {
	name := "TestLoadCertificates"

	for _, tc := range generateTestCases() {
		err := checkLoadCertsFunc(LoadCertificates, tc)
		if err != nil {
			t.Errorf("test %v error: %v", name, err)
		}
	}

} //  End of function  TestLoadCertificates.

// Check validate certificates.
func checkValidateCertsFunc(fn ValidateCertsFunc, tc tlsTestCase) error {
	name := functionName(fn)

	fail := func(unit, msg string, err error) error {
		if err == nil {
			return fmt.Errorf("%v unit %v %v", name, unit, msg)
		}

		return fmt.Errorf("%v error: %w", msg, err)
	}

	err := fn(tc.certpath)

	if tc.ncerts > 0 {
		if err != nil {
			return fail(tc.name, "unexpected", err)
		}

		return nil
	}

	// we didn't get any certs, so expect errors.
	if err != nil {
		return nil
	}

	return fail(tc.name, "expected cert error", err)

} //  End of function  checkValidateCertsFunc.

// Test ValidateCertificate function.
func TestValidateCertificate(t *testing.T) {
	name := "TestLoadCertificates"

	for _, tc := range generateTestCases() {
		err := checkValidateCertsFunc(ValidateCertificates, tc)
		if err != nil {
			t.Errorf("test %v error: %v", name, err)
		}
	}

} //  End of function TestValidateCertificate.

// Test LoadCACerts function.
func TestLoadCACerts(t *testing.T) {
	name := "TestLoadCACerts"

	for _, tc := range generateTestCases() {
		err := checkLoadCertsFunc(LoadCACerts, tc)
		if err != nil {
			t.Errorf("test %v error: %v", name, err)
		}
	}

} //  End of function  TestLoadCACerts.

// Test ValidateCACerts function.
func TestValidateCACerts(t *testing.T) {
	name := "TestValidateCACerts"

	for _, tc := range generateTestCases() {
		err := checkValidateCertsFunc(ValidateCACerts, tc)
		if err != nil {
			t.Errorf("test %v error: %v", name, err)
		}
	}

} //  End of function TestValidateCACerts.

// Test LoadPrivateKeys function.
func TestLoadPrivateKeys(t *testing.T) {
	for _, tc := range generateTestCases() {
		privKeys, err := LoadPrivateKeys(tc.keypath)
		if tc.nkeys > 0 {
			if err != nil {
				t.Errorf("test %v unexpected error: %v",
					tc.name, err)
			}

			nkeys := 0
			if privKeys != nil {
				nkeys = len(privKeys)
			}

			if nkeys != tc.nkeys {
				t.Errorf("test %v expected %v keys got %v",
					tc.name, tc.nkeys, nkeys)
			} else {
				t.Logf("test %v matched %v certs", tc.name,
					tc.nkeys)
			}

			continue
		}

		// we didn't get any keys, so expect errors.
		if err == nil {
			t.Errorf("test %v expected key error", tc.name)
		}
	}

} //  End of function  TestLoadPrivateKeys.

// Test ValidatePrivateKeys function.
func TestValidatePrivateKeys(t *testing.T) {
	for _, tc := range generateTestCases() {
		err := ValidatePrivateKeys(tc.keypath)
		if tc.nkeys > 0 {
			if err != nil {
				t.Errorf("test %v unexpected error: %v",
					tc.name, err)
			}

			continue
		}

		// we didn't get any keys, so expect errors.
		if err == nil {
			t.Errorf("test %v expected key error", tc.name)
		}
	}

} //  End of function  TestValidatePrivateKeys.

// Test LoadCertKeyPair function.
func TestLoadCertKeyPair(t *testing.T) {
	for _, tc := range generateTestCases() {
		cert, err := LoadCertKeyPair(tc.certpath, tc.keypath)
		if !tc.loadErr {
			if err != nil {
				t.Errorf("test %v unexpected error: %v",
					tc.name, err)
			}

			if cert.Leaf == nil {
				t.Errorf("test %v expected a leaf cert",
					tc.name)
			}

			if len(cert.Leaf.Subject.String()) == 0 {
				t.Errorf("test %v expected a certificate",
					tc.name)
			}

			t.Logf("test %v subject: %v", tc.name,
				cert.Leaf.Subject.String())
			continue
		}

		// we expect errors.
		if err == nil {
			t.Errorf("test %v expected an error", tc.name)
		}
	}

} //  End of function  TestLoadCertKeyPair.

// Test ValidateCertKeyPair function.
func TestValidateCertKeyPair(t *testing.T) {
	for _, tc := range generateTestCases() {
		err := ValidateCertKeyPair(tc.certpath, tc.keypath)
		if !tc.loadErr {
			if err != nil {
				t.Errorf("test %v unexpected error: %v",
					tc.name, err)
			}

			continue
		}

		// we expect errors.
		if err == nil {
			t.Errorf("test %v expected an error", tc.name)
		}
	}

} //  End of function  TestValidateCertKeyPair.

// Generate TLS CA certificate specific test cases.
func generateCACertTestCases() []tlsCATestCase {
	return []tlsCATestCase{
		{
			name:      "empty",
			cacerts:   []string{},
			loadError: false,
		},
		{
			name:      "nil",
			cacerts:   nil,
			loadError: false,
		},
		{
			name:      "missing cacert",
			cacerts:   []string{"missing-404-cert.pem"},
			loadError: true,
		},
		{
			name: "missing cacerts",
			cacerts: []string{"d1.pem", "d2.pem", "dev3.pem",
				"s404.pem"},
			loadError: true,
		},
		{
			name:      "single service ca cert",
			cacerts:   []string{getPath(SERVICE, "cacert.pem")},
			loadError: false,
		},
		{
			name:      "single device ca cert",
			cacerts:   []string{getPath(DEVICE, "news-cacert.pem")},
			loadError: false,
		},
		{
			name:      "single bootstrap ca cert",
			cacerts:   []string{getPath(BOOTSTRAP, "bootstrap-cacert.pem")},
			loadError: false,
		},
		{
			name:      "single service ca bundle cert",
			cacerts:   []string{getPath(SERVICE_BUNDLE, "ca.pem")},
			loadError: false,
		},
		{
			name:      "single device ca bundle cert",
			cacerts:   []string{getPath(DEVICE_BUNDLE, "ca-alchemy.pem")},
			loadError: false,
		},
		{
			name:      "single bootstrap ca bundle cert",
			cacerts:   []string{getPath(BOOTSTRAP_BUNDLE, "ca-bootstrap.pem")},
			loadError: false,
		},
		{
			name: "multi ca certs",
			cacerts: []string{getPath(SERVICE, "cacert.pem"),
				getPath(BOOTSTRAP, "bootstrap-cacert.pem"),
				getPath(DEVICE, "alchemy-cacert.pem"),
				getPath(DEVICE, "communique-cacert.pem"),
			},
			loadError: false,
		},
		{
			name: "multi ca certs with a few missing",
			cacerts: []string{getPath(SERVICE, "cacert.pem"),
				"missing-404-cert.pem",
				getPath(BOOTSTRAP, "no-cacert.pem"),
				getPath(DEVICE, "alchemy-cacert.pem"),
				"invalid-cert.pem",
				getPath(DEVICE, "communique-cacert.pem"),
			},
			loadError: true,
		},
		{
			name: "multiple bundled ca certs",
			cacerts: []string{
				getPath(SERVICE_BUNDLE, "ca.pem"),
				getPath(BOOTSTRAP_BUNDLE,
					"ca-bootstrap.pem"),
				getPath(DEVICE_BUNDLE, "ca-alchemy.pem"),
				getPath(DEVICE_BUNDLE, "ca-news.pem"),
				getPath(DEVICE_BUNDLE, "ca-telegraph.pem"),
			},
			loadError: false,
		},
		{
			name: "mixed bundles and single ca certs",
			cacerts: []string{
				getPath(DEVICE, "alchemy-cacert.pem"),
				getPath(DEVICE, "communique-cacert.pem"),
				getPath(DEVICE_BUNDLE,
					"ca-industrial-disease.pem"),
				getPath(DEVICE, "news-cacert.pem"),
				getPath(DEVICE_BUNDLE, "ca-telegraph.pem"),
				getPath(SERVICE, "cacert.pem"),
				getPath(SERVICE_BUNDLE, "ca.pem"),
				getPath(BOOTSTRAP, "bootstrap-cacert.pem"),
				getPath(BOOTSTRAP_BUNDLE,
					"ca-bootstrap.pem"),
			},
		},
		{
			name:      "bad single ca cert",
			cacerts:   []string{getPath(DEVICE, "news-csr.pem")},
			loadError: true,
		},
		{
			name: "bad multiple ca certs",
			cacerts: []string{
				getPath(DEVICE, "alchemy-csr.pem"),
				getPath(DEVICE, "communique-csr.pem"),
				getPath(DEVICE, "news-key.pem"),
			},
			loadError: true,
		},
		{
			name: "mixed good and bad ca certs",
			cacerts: []string{
				getPath(DEVICE, "news-cacert.pem"),
				getPath(DEVICE, "alchemy-csr.pem"),
				getPath(DEVICE_BUNDLE,
					"ca-industrial-disease.pem"),
				getPath(DEVICE, "communique-csr.pem"),
				getPath(BOOTSTRAP, "bootstrap-csr.pem"),
				getPath(SERVICE_BUNDLE, "ca.pem"),
				getPath(SERVICE, "csr.pem"),
				getPath(SERVICE, "key.pem"),
				getPath(SERVICE, "cacert.pem"),
			},
			loadError: true,
		},
	}

} // End of function  generateCACertTestCases.

// Test CreateCACertPool function.
func TestCreateCACertPool(t *testing.T) {
	for _, step := range generateCACertTestCases() {
		pool, err := CreateCACertPool(step.cacerts)
		if !step.loadError {
			if err != nil {
				t.Errorf("test %v unexpected error: %v",
					step.name, err)
			}

			if pool == nil {
				t.Errorf("test %v expected a cert pool",
					step.name)
			}

			continue
		}

		// we expect errors.
		if err == nil {
			t.Errorf("test %v expected create cert pool error",
				step.name)
		}
	}

} //  End of function  TestCreateCACertPool.

// Check device config.
func checkDeviceConfig(tc tlsTestCase, step tlsCATestCase, t *testing.T) {
	name := fmt.Sprintf("test %v + %v", tc.name, step.name)
	cfg, err := DeviceConfig(tc.certpath, tc.keypath, step.cacerts)
	if tc.loadErr || step.loadError {
		// we expect errors.
		if err == nil {
			t.Errorf("%v expected device config error", name)
		}

		return
	}

	if err != nil {
		t.Errorf("%v unexpected error: %v", name, err)
	}

	if cfg == nil {
		t.Errorf("%v expected device config", name)
	}

} // End of function  checkDeviceConfig.

// Test DeviceConfig function.
func TestDeviceConfig(t *testing.T) {
	for _, tc := range generateTestCases() {
		for _, step := range generateCACertTestCases() {
			checkDeviceConfig(tc, step, t)
		}
	}

} //  End of function  TestDeviceConfig.

// Check service config.
func checkServiceConfig(tc tlsTestCase, step tlsCATestCase, t *testing.T) {
	authTypes := []tls.ClientAuthType{
		tls.NoClientCert,
		tls.RequestClientCert,
		tls.RequireAnyClientCert,
		tls.VerifyClientCertIfGiven,
		tls.RequireAndVerifyClientCert,
	}

	for _, clientAuth := range authTypes {
		name := fmt.Sprintf("test %v with %v", tc.name, step.name)
		cfg, err := ServiceConfig(tc.certpath, tc.keypath,
			step.cacerts, clientAuth)
		if tc.loadErr || step.loadError {
			// we expect errors.
			msg := "expected config error"
			if err == nil {
				t.Errorf("%v %v", name, msg)
			}

			continue
		}

		if err != nil {
			t.Errorf("%v unexpected error: %v", name, err)
		}

		if cfg == nil {
			t.Errorf("%v expected config", name)
		}
	}

} //  End of function  checkServiceConfig.

// Test ServiceConfig function.
func TestServiceConfig(t *testing.T) {
	for _, tc := range generateTestCases() {
		for _, step := range generateCACertTestCases() {
			checkServiceConfig(tc, step, t)
		}
	}

} //  End of function  TestServiceConfig.

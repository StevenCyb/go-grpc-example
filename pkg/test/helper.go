package tests

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func GetFreeLocalPort(t *testing.T) int {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer listener.Close()

	addr, ok := listener.Addr().(*net.TCPAddr)
	assert.True(t, ok)
	port := addr.Port

	return port
}

func GenerateAndLoadX509KeyPairWithCA(t *testing.T) (caPrivateKeyPEM, caCertificatePEM, privateKeyPEM, certificatePEM string) {
	t.Helper()

	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key for CA: %v", err)
	}

	caCertTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "CA Root",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caCertificateDER, err := x509.CreateCertificate(rand.Reader, &caCertTemplate, &caCertTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create CA X.509 certificate: %v", err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key for the server: %v", err)
	}

	certTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
	}

	certificateDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, &caCertTemplate, &privateKey.PublicKey, caPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create server X.509 certificate: %v", err)
	}

	caPrivateKeyBytes, err := x509.MarshalPKCS8PrivateKey(caPrivateKey)
	if err != nil {
		t.Fatalf("Failed to marshal CA private key to PKCS#8: %v", err)
	}

	caPrivateKeyPEMBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: caPrivateKeyBytes}
	caCertificatePEMBlock := &pem.Block{Type: "CERTIFICATE", Bytes: caCertificateDER}
	privateKeyPEMBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	certificatePEMBlock := &pem.Block{Type: "CERTIFICATE", Bytes: certificateDER}

	caPrivateKeyPEM = string(pem.EncodeToMemory(caPrivateKeyPEMBlock))
	caCertificatePEM = string(pem.EncodeToMemory(caCertificatePEMBlock))
	privateKeyPEM = string(pem.EncodeToMemory(privateKeyPEMBlock))
	certificatePEM = string(pem.EncodeToMemory(certificatePEMBlock))

	return caPrivateKeyPEM, caCertificatePEM, privateKeyPEM, certificatePEM
}

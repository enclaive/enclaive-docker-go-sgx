package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"os"
	"time"
	"fmt"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello!\n"))
}

func CreateCerts() {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("/tmp/server.crt")
	if err != nil {
		panic(err)
	}
	pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	b, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		panic(err)
	}
	f, err = os.Create("/tmp/server.key")
	if err != nil {
		panic(err)
	}
	pem.Encode(f, &pem.Block{Type: "PRIVATE KEY", Bytes: b})

	fmt.Println("created /tmp/server.key and /tmp/server.crt");
}

func main() {

	CreateCerts()

	http.HandleFunc("/hello", HelloServer)
	fmt.Println("listening on :443");
	err := http.ListenAndServeTLS(":443", "/tmp/server.crt", "/tmp/server.key", nil)
	if err != nil {
		panic(err)
	}
}


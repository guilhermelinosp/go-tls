package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func main() {
	clientCertFile := "keys/client.crt"
	clientKeyFile := "keys/client.key"
	caCertFile := "keys/ca.crt"

	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatalf("Failed to load client cert: %v", err)
	}

	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{clientCert},
				RootCAs:      caCertPool,
			},
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}

	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)
}

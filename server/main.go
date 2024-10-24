package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, TLS!")
	log.Printf("Request received from %s\n", r.RemoteAddr)
}

func main() {
	certFile := "keys/server.crt"
	keyFile := "keys/server.key"
	clientCAFile := "keys/ca.crt"

	caCert, err := os.ReadFile(clientCAFile)
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	server := &http.Server{
		Addr:    ":8443",
		Handler: http.HandlerFunc(helloHandler),
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  caCertPool,
		},
	}

	log.Println("Starting server on https://localhost:8443")
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

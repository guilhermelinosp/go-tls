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
	// Caminhos para os arquivos de certificado e chave
	certFile := "server.crt" // Caminho para o certificado do servidor
	keyFile := "server.key"  // Caminho para a chave privada do servidor
	clientCAFile := "ca.crt" // Caminho para o certificado da CA

	// Carregar o certificado CA
	caCert, err := os.ReadFile(clientCAFile)
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configuração do TLS com verificação de cliente
	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert, // Requer que o cliente apresente um certificado
		ClientCAs:  caCertPool,                     // Usar a CA para verificar os certificados do cliente
	}

	// Criar o servidor HTTP
	server := &http.Server{
		Addr:      ":8443",
		Handler:   http.HandlerFunc(helloHandler),
		TLSConfig: tlsConfig,
	}

	log.Println("Starting server on https://localhost:8443")
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

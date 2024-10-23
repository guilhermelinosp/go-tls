package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func main() {
	// Caminhos para os arquivos de certificado e chave do cliente
	clientCertFile := "client.crt" // Caminho para o certificado do cliente
	clientKeyFile := "client.key"  // Caminho para a chave privada do cliente
	caCertFile := "ca.crt"         // Caminho para o certificado da CA

	// Carregar o certificado do cliente e a chave
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatalf("Failed to load client cert: %v", err)
	}

	// Carregar o certificado CA para verificar o servidor
	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configurar o TLS para usar o certificado do cliente
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert}, // Usar o certificado do cliente
		RootCAs:      caCertPool,                    // Usar a CA para verificar o servidor
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{
		Transport: transport,
	}

	// Fazer a requisição
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)
}

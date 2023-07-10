package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Read the CA certificate
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}

	// Create a certificate pool and add the CA certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	/* Enable this to send client cert to server for mTLS
	// Read the client certificate and private key
	clientCert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		panic(err)
	}*/

	// Create a TLS configuration with the CA certificate pool
	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: false, // Enable certificate verification
		//Certificates:       []tls.Certificate{clientCert},
	}

	// Create an HTTP client with the custom TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Send an HTTPS request to the server
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response:", string(body))
}

package main

import (
	"crypto/tls"
	//"crypto/x509"
	"fmt"
	//"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, client!\n")
	})

	serverCert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}

	/* Enable this to validate client cert
	// Load the CA certificate
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}

	// Create a certificate pool and add the CA certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	*/

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		//ClientCAs:    caCertPool,
		//ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{serverCert},
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server listening on https://localhost:8443")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

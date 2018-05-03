package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	// strip HTTP port but do not add back HTTPS port, on the assumption that we're binding to
	// a port >1024 but HTTPS is being redirected to the standard port 443
	host := strings.Split(r.Host, ":")

	http.Redirect(w, r, "https://"+host[0]+r.RequestURI, http.StatusFound)
}

func main() {
	// options with their defaults
	options := struct {
		host                string
		port                uint
		wwwRoot             string
		sslEnable           bool
		sslLetsEncrypt      bool
		sslPort             uint
		redirectHTTP        bool
		sslCert             string
		sslKey              string
		sslLetsEncryptCerts string
	}{
		host:                "",
		port:                8000,
		wwwRoot:             ".",
		sslEnable:           false,
		sslLetsEncrypt:      false,
		sslPort:             8043,
		redirectHTTP:        true,
		sslCert:             "",
		sslKey:              "",
		sslLetsEncryptCerts: "certs",
	}

	// parse command line options
	flag.StringVar(&options.host, "host", options.host, "Hostname or IP to bind to (empty string for all IPs on this host)")
	flag.UintVar(&options.port, "p", options.port, "Port to bind to")
	flag.StringVar(&options.wwwRoot, "www", options.wwwRoot, "Directory to serve")
	flag.BoolVar(&options.sslEnable, "sslenable", options.sslEnable, "Enable TLS/SSL (HTTPS)")
	flag.BoolVar(&options.sslLetsEncrypt, "sslletsencrypt", options.sslLetsEncrypt, "Use Let's Encrypt for TLS/SSL certificate issuance and renewal")
	flag.UintVar(&options.sslPort, "sslport", options.sslPort, "TLS/SSL port to listen on")
	flag.BoolVar(&options.redirectHTTP, "redirecthttp", options.redirectHTTP, "Redirect HTTP request to HTTPS")
	flag.StringVar(&options.sslCert, "sslcert", options.sslCert, "TLS/SSL Certificate file")
	flag.StringVar(&options.sslKey, "sslkey", options.sslKey, "TLS/SSL Key file")
	flag.StringVar(&options.sslLetsEncryptCerts, "sslletsencryptcerts", options.sslLetsEncryptCerts, "Let's Encrypt certificate cache directory")

	flag.Parse()

	bindTo := fmt.Sprintf("%v:%d", options.host, options.port)
	bindToSSL := fmt.Sprintf("%v:%d", options.host, options.sslPort)

	http.Handle("/", http.FileServer(http.Dir(options.wwwRoot)))

	server := &http.Server{}

	if options.sslEnable {
		server.Addr = bindToSSL

		if options.sslLetsEncrypt {
			// serve via HTTPS with certificates auto-renewed via Let's Encrypt
			certManager := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: autocert.HostWhitelist(options.host),
				Cache:      autocert.DirCache(options.sslLetsEncryptCerts), // folder for storing certificates
			}

			// optionally redirect all traffic on HTTP port to HTTPS (SSL/TLS) port
			if options.redirectHTTP {
				fmt.Printf("Redirecting all HTTP traffic sent to '%v' to HTTPS at '%v'\n", bindTo, bindToSSL)

				go func() {
					// wrap redirect handler in certManager's HTTP handler to support http-01 challenges
					if err := http.ListenAndServe(bindTo, certManager.HTTPHandler(http.HandlerFunc(redirectTLS))); err != nil {
						log.Fatalf("ListenAndServe error: %v", err)
					}
				}()
			} else {
				go func() {
					// launch server on port :80 anyway with certManager's HTTP handler to support http-01 challenges
					if err := http.ListenAndServe(bindTo, certManager.HTTPHandler(nil)); err != nil {
						log.Fatalf("ListenAndServe error: %v", err)
					}
				}()
			}

			server.TLSConfig = &tls.Config{
				GetCertificate: certManager.GetCertificate,
			}

			fmt.Printf("Let's Encrypt TLS/SSL web server started at '%v' serving directory '%v'\n", bindToSSL, options.wwwRoot)

			log.Fatal(server.ListenAndServeTLS("", "")) // key and cert are comming from Let's Encrypt
		} else {
			// serve via HTTPS with static certficate

			// optionally redirect all traffic on HTTP port to HTTPS (SSL/TLS) port
			if options.redirectHTTP {
				fmt.Printf("Redirecting all HTTP traffic sent to '%v' to HTTPS at '%v'\n", bindTo, bindToSSL)
	
				go func() {
					if err := http.ListenAndServe(bindTo, http.HandlerFunc(redirectTLS)); err != nil {
						log.Fatalf("ListenAndServe error: %v", err)
					}
				}()
			}

			fmt.Printf("TLS/SSL web server started at '%v' serving directory '%v'\n", bindToSSL, options.wwwRoot)

			log.Fatal(server.ListenAndServeTLS(options.sslCert, options.sslKey))
		}
	} else {
		// serve via HTTP
		server.Addr = bindTo

		fmt.Printf("Serving directory '%v' via HTTP on '%v'\n", options.wwwRoot, bindTo)

		log.Fatal(server.ListenAndServe())
	}
}

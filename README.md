# kantand
Very simple http server written in Go for local testing purposes

## Overview
This is intended for developers who are testing a static website on their local systems (localhost, 127.0.0.1, ::1, etc.). By default it will bind to port `8000` on all available local IP addresses and serve the local directory recursively. **This is convenient for developers who may want to access the website from another system, but is a glaring security hole if you don't want others to be able to access it.**

The name `kantand` was chosen to be easy enough to type, and because in Japanese it means "easy daemon" (かんたんd).

## Usage
To install:

```sh
$ go get github.com/shoenseiwaso/kantand
```

To run, assuming `$GOPATH/bin` is in your `$PATH`:

```sh
$ kantand
Serving directory '.' via HTTP on ':8000'.
```

To quit, press `CTRL+C`.

A safer option would be to bind only to localhost:

```sh
$ kantand -host localhost
Serving directory '.' via HTTP on 'localhost:8000'.
```

With TLS/SSL using Let's Encrypt (requires an internet-visible hostname):

```sh
$ kantand -host test.example.com -sslenable -sslletsencrypt
```

Optional parameters:

```sh
$ kantand -h
Usage of ./kantand:
  -host string
    	Hostname or IP to bind to (empty string for all IPs on this host)
  -p uint
    	Port to bind to (default 8000)
  -redirecthttp
    	Redirect HTTP request to HTTPS (default true)
  -sslcert string
    	TLS/SSL Certificate file
  -sslenable
    	Enable TLS/SSL (HTTPS)
  -sslkey string
    	TLS/SSL Key file
  -sslletsencrypt
    	Use Let's Encrypt for TLS/SSL certificate issuance and renewal
  -sslport uint
    	TLS/SSL port to listen on (default 8043)
  -www string
    	Directory to serve (default ".")
```

## Licence
MIT

## Author
Jeff Wentworth, co-founder of [Curvegrid](http://curvegrid.com), a blockchain tooling company.
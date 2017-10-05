# kantand
Very simple http server written in Go for local testing purposes

## Overview
This is intended for developers who are testing a static website on their local systems (localhost, 127.0.0.1, ::1, etc.). By default it will bind to port `8000` on `127.0.0.1` and serve the local directory recursively.

## Usage
To install:

```sh
$ go install github.com/shoenseiwaso/kantand
```

To run, assuming `$GOPATH/bin` is in your `$PATH`:

```sh
$ kantand
```

## Licence
MIT

## Author
Jeff Wentworth, co-founder of [Curvegrid](http://curvegrid.com), a blockchain tooling company.
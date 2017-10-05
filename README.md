# kantand
Very simple http server written in Go for local testing purposes

## Overview
This is intended for developers who are testing a static website on their local systems (localhost, 127.0.0.1, ::1, etc.). By default it will bind to port `8000` on `127.0.0.1` and serve the local directory recursively.

The name `kantand` was chosen to be easy enough to type, and because in Japanese it means "easy daemon" (かんたんd).

## Usage
To install:

```sh
$ go install github.com/shoenseiwaso/kantand
```

To run, assuming `$GOPATH/bin` is in your `$PATH`:

```sh
$ kantand
Serving directory '.' via HTTP on 'localhost:8000'.
```

To quit, press `CTRL+C`.

Optional parameters:

```sh
$ kantand -h
Usage of kantand:
  -bind string
    	Host and port to bind to (default ":8000")
  -h	Display help text
  -www string
    	Directory to serve (default ".")
```

## Licence
MIT

## Author
Jeff Wentworth, co-founder of [Curvegrid](http://curvegrid.com), a blockchain tooling company.
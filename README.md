# HTTP Hands-on

Sample codes for Software Design September 2017.


## Installation

To install the library and command line program, use the following:

```
go get -u github.com/Ladicle/http-handson
```

## Usage

Run HTTP server:

```
$ http-handson
2017/07/11 06:46:58 Access to 127.0.0.1:7070
```

Show command usage:

```
$ http-handson --help
Usage of ./http-handson:
  -h string
        server host name (default "127.0.0.1")
  -p int
        server port number (default 7070) 
```


## Reconvert views

Views are converted to binary data by `go-bindata`.
If you want to reconvert view, to run the the following command.

```
$ go get -u github.com/jteeuwen/go-bindata/...
$ go-bindata -o view.go views
```

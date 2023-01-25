# TCP Mirror

`tcpmirror` duplicates TCP traffic received on a port to more than one destination ports.
One of the destination is the primary server (port) which responds
to the incoming TCP traffic. Other streams are duplicate streams and
any response from the respective servers is discarded.

## Why

This is helpful in the following scenarios:

- To test a Dev/QA/secondary server with the same requests, traffic
and load that the primary production server handles.

- To do performance testing of candidate servers against the existing
production server.

- To re-write the server in another language and make sure that the
new server responds the same as existing server for same requests.

## Install

### Download from the releases pages

Download pre-built binary from the release page.

### Use `go get`

If you have `golang` tools installed, you can download and build the source code
locally as follows:
```
$ go get github.com/codeexpress/tcpmirror
```
The `tcpmirror` binary is now available in your `$GOPATH/bin` directory

### Compile from source

```
$ git clone https://github.com/codeexpress/tcpmirror.git
$ cd tcpmirror; go run tcpmirror.go
```

## Sample use case

```Usage:   $ tcpmirror -l <listen_addr> -p <primary_addr> -m <mirror_addrs>```

Example:
```
 tcpmirror -l localhost:8080 -p localhost:9090 -m localhost:9091,localhost:9091
```

Flags:
 * **l** - listen address
 * **p** - primary traffic address **(can not be null and must work)**
 * ~~m~~ - list of mirror traffic addresses (can be empty)

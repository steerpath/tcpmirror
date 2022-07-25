package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	Version = "1.0.0"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func main() {

	flag.Usage = Usage
	flag.Parse()

	mirrorAddrs := []string{}

	//mirrorAddrs := strings.Split(*mirrorPtr, ",")
	fmt.Println(mirrorAddrs)

	listen := getEnvOrDefault("LISTEN", "localhost:8080")
	primary := getEnvOrDefault("PRIMARY", "localhost:9090")
	mirrors := getEnvOrDefault("MIRRORS", "")

	listenPtr := &listen
	primaryPtr := &primary
	if mirrors != "" {
		mirrorAddrs = strings.Split(mirrors, ",")
	}

	fmt.Println("Hello TCP Mirror")
	fmt.Printf("Listening on                    %s\n", *listenPtr)
	fmt.Printf("Connecting in primary mode to   %s\n", *primaryPtr)
	if len(mirrorAddrs) > 0 {
		fmt.Printf("Connecting in mirror mode to   %s\n", mirrorAddrs)
	}

	l, err := net.Listen("tcp", *listenPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening: ", err.Error())
		os.Exit(1)
	}

	for {
		in, _ := l.Accept()
		fmt.Printf("Incoming connection from %s\n", in)

		p, err := net.Dial("tcp", *primaryPtr)
		if err != nil {
			fmt.Println("Error connecting to primary: ", err.Error())
			continue
		}

		ws := make([]io.Writer, len(mirrorAddrs))
		if len(mirrorAddrs) > 0 {

			for i, mirrorAddr := range mirrorAddrs {
				m, err := net.Dial("tcp", mirrorAddr)
				fmt.Println("Sending data to mirror: ", mirrorAddr)
				if err != nil {
					fmt.Println("Error connecting to the mirror address: ", err.Error())
					continue
				}
				ws[i] = m
			}

		}
		ws = append(ws, p) // add primary
		mw := io.MultiWriter(ws...)
		go io.Copy(mw, in)

		go io.Copy(in, p)

		// fmt.Println("After accept")
		// fmt.Printf("mw = %v\nin = %v\n", mw, in)
		// fmt.Printf("Num goroutines: %d \n", runtime.NumGoroutine())
	}
	// Close the listener when application closes
}

func Usage() {
	fmt.Fprintf(os.Stderr, "tcpmirror version %s\n", Version)
	fmt.Fprintf(os.Stderr, "Usage:   $ tcpmirror -l <listen_addr> -p <primary_addr> -m <mirror_addrs\n")
	fmt.Fprintf(os.Stderr, "Example: $ tcpmirror -l localhost:8080 -p localhost:9090 -m localhost:9091,localhost:9091 \n")
	fmt.Fprintf(os.Stderr, "-----------------------\nFlags:\n")
	flag.PrintDefaults()
}

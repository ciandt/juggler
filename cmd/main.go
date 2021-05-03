package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"com.ciandt.juggler/internal"
)

var s5Address string
var s5User string
var s5Password string

var errInvalidSocks5Address = errors.New("Invalid SOCKS5 address")

const (
	help = `Usage: juggler [proxy_config] [-p local_port] outbound_address
	
Example: SOCKS5 proxy running on localhost:3500 requesting 
server on https://my-protected-api.com:9000 being accessed
on localhost:12345

juggler -port 12345 -s5.address=localhost:9000 -s5.user=socks_user -s5.password=socks_pass https://my-protected-api.com:9000
	
`
)

func main() {

	var port = flag.Int("port", 9001, "Local address destination port")
	flag.StringVar(&s5Address, "s5.address", "", "SOCKS5 address")
	flag.StringVar(&s5User, "s5.user", "", "SOCKS5 authentication user")
	flag.StringVar(&s5Password, "s5.password", "", "SOCKS5 authenticaion password")

	flag.Parse()

	outbound := flag.Arg(0)
	if outbound == "" {
		fmt.Println("Mandatory outbound address not provided")
		printUsage()
		os.Exit(1)
	}

	ob := internal.Outbound{
		Address: outbound,
	}

	srv, err := internal.NewServer(*port, 2*time.Minute, 2*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	s5cfg, err := parseSocks5Config()

	if err == nil {
		log.Println("SOCKS5 config found, using it to proxy requests")
		srv.ProxySocks5(ob, s5cfg)
	}

	srv.ProxyDefaultTCP(ob)
}

func parseSocks5Config() (internal.Socks5Config, error) {

	if s5Address == "" {
		log.Println("Socks5 address invalid or not found. Fallbacking to another proxy method")
		return internal.Socks5Config{}, errInvalidSocks5Address
	}

	return internal.Socks5Config{
		Address:  s5Address,
		User:     s5User,
		Password: s5Password,
	}, nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, help)
	flag.PrintDefaults()
}

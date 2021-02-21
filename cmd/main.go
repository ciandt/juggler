package main

import (
	"flag"

	"com.ciandt.juggler/internal"
)

func main() {

	var outbound = flag.String("outbound", "localhost:5000", "Outbound destination server")
	var s5Address = flag.String("s5.adress", "localhost:3333", "socks5 address")
	var s5User = flag.String("s5.user", "", "socks5 authentication user")
	var s5Password = flag.String("s5.password", "", "socks5 authenticaion password")

	flag.Parse()

	ob := internal.Outbound{
		Address: *outbound,
	}

	socks5 := internal.Socks5Config{
		Address:  *s5Address,
		User:     *s5User,
		Password: *s5Password,
	}

	internal.ProxySocks5(ob, socks5)
}

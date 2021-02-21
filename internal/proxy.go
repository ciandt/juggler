package internal

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"golang.org/x/net/proxy"
)

type Socks5Config struct {
	Address  string
	User     string
	Password string
}

type Outbound struct {
	Address string
}

func ProxySocks5(outbound Outbound, cfg Socks5Config) {
	log.Printf("Creating Socks5 proxy for %s via socks5://%s",
		outbound.Address,
		cfg.Address)
	url, err := url.Parse(outbound.Address)

	if err != nil {
		log.Fatal(err)
	}

	rp := httputil.NewSingleHostReverseProxy(url)
	rp.Transport = createSocks5Transport(cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func createSocks5Transport(cfg Socks5Config) *http.Transport {
	auth := &proxy.Auth{}
	if cfg.User != "" || cfg.Password != "" {
		auth = &proxy.Auth{
			User:     cfg.User,
			Password: cfg.Password,
		}
	}

	dialSocksProxy, err := proxy.SOCKS5("tcp", cfg.Address, auth, proxy.Direct)
	if err != nil {
		log.Println("Error connecting to socks5 proxy:", err)
	}
	return &http.Transport{Dial: dialSocksProxy.Dial}
}

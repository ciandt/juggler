package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

//Socks5Config holds configuration for a possible socks5 connection proxy
//It's an optional configuration
type Socks5Config struct {
	Address  string
	User     string
	Password string
}

//Outbound holds all configuration related to the actual
//service to be called. It's a mandatory structure
type Outbound struct {
	Address string
}

//Server is the instance to expose the local endpoint to proxy
//the request to the outbound service
type Server struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//NewServer creates a new server instance. The port should be a valid
//int used to host the local listener port (1-65535)
func NewServer(
	port int,
	readTimeout time.Duration,
	writeTimeout time.Duration,
) (*Server, error) {

	if port <= 0 || port > 65535 {
		return &Server{}, errors.New("Invalid port number")
	}

	add := fmt.Sprintf(":%d", port)

	return &Server{
		Address:      add,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}, nil

}

//ProxySocks5 does a proxy using a socks5 custom transport
func (srv *Server) ProxySocks5(outbound Outbound,
	cfg Socks5Config) {
	log.Printf("Creating Socks5 proxy for %s via SOCKS5://%s",
		outbound.Address,
		cfg.Address)

	transport := createSocks5Transport(cfg)

	srv.doProxy(outbound, transport)
}

func (srv *Server) doProxy(
	outbound Outbound,
	transport *http.Transport) {

	url, err := url.Parse(outbound.Address)

	if err != nil {
		log.Fatal(err)
	}

	rp := httputil.NewSingleHostReverseProxy(url)
	rp.Transport = transport

	outHost := dropSchemaAndPort(outbound.Address)
	log.Printf("Replacing the original request host \"localhost\" by the provided outbound \"%s\"", outHost)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Host = outHost
		logRequest(r)
		rp.ServeHTTP(w, r)
	})

	log.Printf("Forwarding all requests made to %s to its proxy", srv.Address)

	log.Fatal(http.ListenAndServe(srv.Address, nil))
}

func dropSchemaAndPort(address string) string {
	uri, err := url.Parse(address)
	if err != nil {
		log.Fatalf("Error while parsing address %s", address)
	}
	return uri.Hostname()
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
		log.Fatalln("Error connecting to socks5 proxy:", err)
	}
	return &http.Transport{Dial: dialSocksProxy.Dial}
}

func logRequest(r *http.Request) {
	log.Printf("Forwarding request made to path %s", r.URL)
}

package main

import "testing"

func TestShouldReturnErrorOnEmptySocks5Address(t *testing.T) {

	s5Address = ""

	_, err := parseSocks5Config()
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestShouldReturnConfigOnValidSocks5Address(t *testing.T) {

	s5Address = "localhost:1234"
	s5User = "admin"
	s5Password = "admin-pass"

	cfg, err := parseSocks5Config()
	if err != nil {
		t.Error("Unexpected parse error", err)
	}

	if cfg.Address != s5Address {
		t.Errorf("Expected %s socks5 address, got %s", s5Address, cfg.Address)
	}

	if cfg.User != s5User {
		t.Errorf("Expected %s socks5 user, got %s", s5User, cfg.User)
	}

	if cfg.Password != s5Password {
		t.Errorf("Expected %s socks5 password, got %s", s5Password, cfg.Password)
	}
}

package internal

import "testing"

func TestShouldDropSchemaAndPortNumber(t *testing.T) {
	hn := dropSchemaAndPort("https://my-hostname.com:9091")

	if hn != "my-hostname.com" {
		t.Errorf("Expected hostname %s, got: %s ", "my-hostname.com", hn)
	}
}

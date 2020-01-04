package sitehash

import (
	"net/url"
	"strings"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
)

func TestFingerprint(t *testing.T) {
	testURLs := []string{
		"http://google.com",
		"http://neverssl.com",
	}
	for _, u := range testURLs {
		t.Run(strings.TrimPrefix(u, "http://"), func(t *testing.T) {
			up, err := url.Parse(u)
			if err != nil {
				t.Fatal(err)
			}
			d, err := Fingerprint(up)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, d)
		})
	}
}

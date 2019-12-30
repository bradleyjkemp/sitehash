package sitehash

import (
	"net"
	"net/http"
	"net/url"
	"sort"
)

type Digest struct {
	Registered  bool
	Nameservers []string
	Status      string
	Headers     []string
}

func Fingerprint(url *url.URL) (Digest, error) {
	d := Digest{}

	var err error
	d.Nameservers, d.Registered, err = getNameservers(url.Hostname())
	if err != nil {
		return Digest{}, err
	}

	if !d.Registered {
		return d, nil
	}

	d.Status, d.Headers, err = getHeaders(url)
	if err != nil {
		return Digest{}, err
	}

	return d, nil
}

func getNameservers(host string) (nameservers []string, exists bool, err error) {
	ns, err := net.LookupNS(host)
	if err, ok := err.(*net.DNSError); ok && err.IsNotFound {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	nameservers = make([]string, 0, len(ns))
	for _, s := range ns {
		nameservers = append(nameservers, s.Host)
	}
	sort.Slice(nameservers, func(i, j int) bool {
		return nameservers[i] < nameservers[j]
	})

	return nameservers, true, nil
}

func getHeaders(url *url.URL) (string, []string, error) {
	resp, err := http.Head(url.String())
	if err != nil {
		return "", nil, err
	}

	headers := make([]string, 0, len(resp.Header))
	for header := range resp.Header {
		headers = append(headers, header)
	}
	sort.Slice(headers, func(i, j int) bool {
		return headers[i] < headers[j]
	})

	return resp.Status, headers, nil
}

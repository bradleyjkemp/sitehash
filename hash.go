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
		return d, err
	}

	if !d.Registered {
		return d, nil
	}

	d.Status, d.Headers, err = getHeaders(url)
	if err != nil {
		return d, err
	}

	return d, nil
}

func getNameservers(host string) (nameservers []string, exists bool, err error) {
	_, ipErr := net.LookupIP(host)
	ns, nsErr := net.LookupNS(host)
	if ipErr != nil && nsErr != nil {
		// both lookups failed so this is likely legitimately not registered
		ipDNSErr, ipDNSOk := ipErr.(*net.DNSError)
		nsDNSErr, nsDNSOk := nsErr.(*net.DNSError)
		if ipDNSOk && ipDNSErr.IsNotFound || nsDNSOk && nsDNSErr.IsNotFound {
			return nil, false, nil
		}
		return nil, false, nsDNSErr
	}

	// If either succeeded, we know the domain exists even if the nameserver lookup failed
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
	// Use GET instead of HEAD to maximise compatibility
	resp, err := http.Get(url.String())
	if err != nil {
		return "", nil, err
	}
	// Discard body
	resp.Body.Close()

	headers := make([]string, 0, len(resp.Header))
	for header := range resp.Header {
		headers = append(headers, header)
	}
	sort.Slice(headers, func(i, j int) bool {
		return headers[i] < headers[j]
	})

	return resp.Status, headers, nil
}

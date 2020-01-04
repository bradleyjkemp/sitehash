package sitehash

import (
	"net/url"
)

type Digest struct {
	Registered  bool
	Hosted      bool
	Nameservers []string
	Status      string
	Headers     []string
}

func Fingerprint(url *url.URL) (d Digest, err error) {
	d.Registered, err = isDomainRegistered(url.Hostname())
	if err != nil {
		return d, err
	}

	if !d.Registered {
		// Domain isn't registered so nothing more we can do
		return d, nil
	}

	d.Nameservers, err = getNameservers(url.Hostname())
	if err != nil {
		return d, err
	}

	d.Hosted, err = isDomainHosted(url.Hostname())
	if err != nil {
		return d, err
	}

	if d.Hosted {
		// Domain is hosted somewhere so try to fetch the given URL
		d.Status, d.Headers, err = getHeaders(url)
		if err != nil {
			return d, err
		}
	}

	return d, nil
}

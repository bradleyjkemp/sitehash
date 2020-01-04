package sitehash

import (
	"net"
	"sort"
)

func isDomainRegistered(host string) (bool, error) {
	_, ipErr := net.LookupIP(host)
	_, nsErr := net.LookupNS(host)
	if ipErr != nil && nsErr != nil {
		// both lookups failed so this is likely legitimately not registered
		ipDNSErr, ipDNSOk := ipErr.(*net.DNSError)
		nsDNSErr, nsDNSOk := nsErr.(*net.DNSError)
		if ipDNSOk && ipDNSErr.IsNotFound || nsDNSOk && nsDNSErr.IsNotFound {
			return false, nil
		}
		return false, ipErr
	}

	// At least one lookup succeeded so this must be registered
	return true, nil
}

func isDomainHosted(host string) (bool, error) {
	_, err := net.LookupIP(host)
	if err == nil {
		return true, nil
	}

	switch err := err.(type) {
	case *net.DNSError:
		if err.IsNotFound {
			return false, nil
		}
		return false, err
	default:
		return false, err
	}
}

func getNameservers(host string) (nameservers []string, err error) {
	ns, err := net.LookupNS(host)
	if err == nil {
		// If either succeeded, we know the domain exists even if the nameserver lookup failed
		nameservers = make([]string, 0, len(ns))
		for _, s := range ns {
			nameservers = append(nameservers, s.Host)
		}
		sort.Slice(nameservers, func(i, j int) bool {
			return nameservers[i] < nameservers[j]
		})
		return nameservers, nil
	}

	switch err := err.(type) {
	case *net.DNSError:
		if err.IsNotFound {
			return nil, nil
		}
		return nil, err
	default:
		return nil, err
	}
}

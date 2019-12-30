# sitehash

A small library to help track large scale changes in domains e.g.
* An un-registered domain becomes registered
* A registered domain starts hosting content
* A domain stops hosting content

This is mainly designed for easily tracking the lifecycle of phishing/malware domains.

### Fingerprinting Strategy

A domain/url's fingerprint is built from:
* Whether or not the domain is registered. This detects when:
    * a previously non-existent domain of interest is registered.
    * an active domain is taken down by the domain registrar.
* What the domain's nameservers are.
    * This usually detects when the domain is sold (as the nameservers change from a domain parking service).
* The HTTP headers returned by the server. All header values are ignored (as random session IDs would cause too much noise).
    * This detects large changes in the content being hosted e.g. a phishing kit is uploaded so now the `Set-Cookie` header is being returned.

### Usage

```go
u, _ := url.Parse("http://example.com")
digest := sitehash.Fingerprint(u)

if !reflect.DeepEqual(digest, previousDigest) {
    fmt.Println("example.com changed!")
}
```

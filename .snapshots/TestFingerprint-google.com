(sitehash.Digest) {
  Registered: (bool) true,
  Hosted: (bool) true,
  Nameservers: ([]string) (len=4) {
    (string) (len=15) "ns1.google.com.",
    (string) (len=15) "ns2.google.com.",
    (string) (len=15) "ns3.google.com.",
    (string) (len=15) "ns4.google.com."
  },
  Status: (string) (len=6) "200 OK",
  Headers: ([]string) (len=9) {
    (string) (len=13) "Cache-Control",
    (string) (len=12) "Content-Type",
    (string) (len=4) "Date",
    (string) (len=7) "Expires",
    (string) (len=3) "P3p",
    (string) (len=6) "Server",
    (string) (len=10) "Set-Cookie",
    (string) (len=15) "X-Frame-Options",
    (string) (len=16) "X-Xss-Protection"
  }
}

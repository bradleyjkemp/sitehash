(sitehash.Digest) {
  Registered: (bool) true,
  Hosted: (bool) true,
  Nameservers: ([]string) (len=4) {
    (string) (len=22) "ns-1413.awsdns-48.org.",
    (string) (len=24) "ns-1716.awsdns-22.co.uk.",
    (string) (len=21) "ns-320.awsdns-40.com.",
    (string) (len=21) "ns-570.awsdns-07.net."
  },
  Status: (string) (len=6) "200 OK",
  Headers: ([]string) (len=11) {
    (string) (len=3) "Age",
    (string) (len=10) "Connection",
    (string) (len=12) "Content-Type",
    (string) (len=4) "Date",
    (string) (len=13) "Last-Modified",
    (string) (len=6) "Server",
    (string) (len=4) "Vary",
    (string) (len=3) "Via",
    (string) (len=11) "X-Amz-Cf-Id",
    (string) (len=12) "X-Amz-Cf-Pop",
    (string) (len=7) "X-Cache"
  }
}

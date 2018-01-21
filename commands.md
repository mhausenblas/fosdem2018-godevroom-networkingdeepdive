# Commands

```
$ go list net/...

$ go doc net/http | head

$ GODEBUG=netdns=cgo go run code/dns/main.go mh9laptop.local
[192.168.0.17 192.168.99.1 192.168.33.1 fe80::102c:9da1:a044:49f0%en0]

$ GODEBUG=netdns=go go run code/dns/main.go mh9laptop.local
panic: lookup mh9laptop.local on 89.101.160.5:53: no such host

goroutine 1 [running]:
main.main()
        /Users/mhausenblas/Dropbox/dev/work/src/github.com/mhausenblas/fosdem2018-godevroom-networkingdeepdive/code/dns/main.go:14 +0xfd
exit status 2
```

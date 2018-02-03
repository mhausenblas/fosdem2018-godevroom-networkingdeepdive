# Commands

## The net package—overview

```
$ go list net/...

$ go doc net/http | head
```

## The net package—curl example

```
$ go run code/callapi/main.go http://mhausenblas.info 100
```

## The net package—appserver

```
$ go run code/app-server/main.go

$ http localhost:9876/about
```

## The net package—RPC client/server example

```
$ go run code/rpc/server/main.go

$ go run code/rpc/client/main.go localhost
```

## Socket-level programming—example

```
$ cat /etc/services
```

```
$ go run code/connection/main.go localhost:3999
```

## DNS—example

```
$ GODEBUG=netdns=cgo go run code/dns/main.go mh9laptop.local
[192.168.0.17 192.168.99.1 192.168.33.1 fe80::102c:9da1:a044:49f0%en0]

$ GODEBUG=netdns=go go run code/dns/main.go mh9laptop.local
panic: lookup mh9laptop.local on 89.101.160.5:53: no such host

goroutine 1 [running]:
main.main()
        /Users/mhausenblas/Dropbox/dev/work/src/github.com/mhausenblas/fosdem2018-godevroom-networkingdeepdive/code/dns/main.go:14 +0xfd
exit status 2
```

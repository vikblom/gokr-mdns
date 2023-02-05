# gokr-mdns

Advertise `gokrazy.local` through mDNS from rasp pi set up through gokrazy.
Useful since for some reason I cannot reach gokrazy pi's just through their hostname.

## Setup
```
gokr-packer github.com/vikblom/gokr-mdns
```
populate `GOKRAZY_UPDATE` to update existing rasp pi on the network.

Use local modifications (fish shell expansion)
```
cd builddir/github.com/vikblom/gokr-mdns/
go mod edit -replace github.com/vikblom/gokr-mdns=(realpath ../../../..)
```

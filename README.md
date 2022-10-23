
1497	129.785455530	fe80::e65f:1ff:fe9a:a1e3	fe80::710c:332d:bd7f:7ff8	MDNS	105	Standard query response 0x0000 A 192.168.0.12

The answer is coming from ipv6 (a1e3)

https://unix.stackexchange.com/questions/241915/firewalld-accept-response-to-multicast-dns-query-from-ephemeral-port


ip route add 224.0.0.251 dev eth0

nc -vv -l -p 5353 -u -s 224.0.0.251
nc -vv -u 224.0.0.251 5353

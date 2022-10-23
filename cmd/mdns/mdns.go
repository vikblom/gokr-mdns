package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/mdns"
)

// func handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello\n")
// }

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// go http.ListenAndServe(":80", http.HandlerFunc(handle))

	// Setup our service export
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	info := []string{"gokrazy mdns"}
	service, err := mdns.NewMDNSService(host, "_gokrazy._tcp", "", "", 80, nil, info)
	if err != nil {
		log.Fatal(err)
	}

	// iface, err := net.InterfaceByName("wlp3s0")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var iface *net.Interface

	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service, Iface: iface, LogEmptyResponses: true})
	if err != nil {
		log.Fatal(err)
	}
	defer server.Shutdown()

	<-ctx.Done()
}

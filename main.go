package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/brutella/dnssd"
	dnssdlog "github.com/brutella/dnssd/log"
	"github.com/gokrazy/gokrazy"
)

func main() {
	dnssdlog.Debug.Enable()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Wait until network interfaces have a chance to work.
	gokrazy.WaitForClock()

	resp, err := dnssd.NewResponder()
	if err != nil {
		log.Fatalf("new responder: %s", err)
	}
	// Just need the A record.
	_, err = resp.Add(dnssd.Service{
		Domain: "local",
		Host:   "gokrazy",
		Ifaces: []string{"eth0", "wlan0"},
	})
	if err != nil {
		log.Fatalf("add service: %s", err)
	}
	err = resp.Respond(ctx)
	if err != nil {
		log.Fatalf("respond: %s", err)
	}
}

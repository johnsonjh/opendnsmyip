// Copyright (c) 2020 Jeffrey H. Johnson.
// Copyright (c) 2020 Gridfinity, LLC.
// Copyright (c) 2020 Dmitriy Kuznetsov.
// Copyright (c) 2016 James Polera.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package opendnsmyip_test

import (
	"testing"

	myip "go.gridfinity.dev/opendnsmyip"
	u "go.gridfinity.dev/leaktestfe"
	"github.com/miekg/dns"
)

func TestGetMyIPAddress(t *testing.T) {
	defer u.LeakPlug(
		t,
	)
	_, err := myip.GetMyIP()
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestGenerateClientError(t *testing.T) {
	defer u.LeakPlug(
		t,
	)
	config := dns.ClientConfig{Servers: []string{"0.0.0.0"}, Port: "53"}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("myip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	_, err := myip.MyIPDNSLookup(config, dnsClient, message)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGenerateLookupError(t *testing.T) {
	defer u.LeakPlug(
		t,
	)
	config := dns.ClientConfig{
		Servers: []string{"208.67.220.220", "208.67.222.222"},
		Port:    "53",
	}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("notmyip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	_, err := myip.MyIPDNSLookup(config, dnsClient, message)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

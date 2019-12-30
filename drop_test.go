package drop

import (
	"context"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
)

func TestDrop(t *testing.T) {
	drop := Drop{Next: test.ErrorHandler(), qtype: dns.TypeAAAA}
	ctx := context.Background()

	req := new(dns.Msg)
	req.SetQuestion("example.org.", dns.TypeAAAA)

	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	drop.ServeDNS(ctx, rec, req)
	if rec.Rcode != 4 {
		t.Errorf("Expected 0 as Rcode, got %d", rec.Rcode)
	}

	req.SetQuestion("example.org.", dns.TypeA)
	drop.ServeDNS(ctx, rec, req)
	if rec.Rcode == 0 {
		t.Errorf("0 not Expected as Rcode")
	}

}

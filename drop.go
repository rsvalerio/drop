package drop

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("drop")

// Drop struct
type Drop struct {
	Next  plugin.Handler
	qtype uint16
}

// ServeDNS implements the plugin.Handler interface. This method gets called when drop is used
// in a Server.
func (d Drop) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debug("Dropping response")
	state := request.Request{W: w, Req: r}

	m := new(dns.Msg)
	m.SetReply(r)

	if state.QType() == d.qtype {
		log.Debug("Received response")
		dropCount.WithLabelValues(metrics.WithServer(ctx)).Inc()
		m.Rcode = dns.RcodeNotImplemented
		w.WriteMsg(m)
		return 0, nil
	}

	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d Drop) Name() string { return "drop" }
